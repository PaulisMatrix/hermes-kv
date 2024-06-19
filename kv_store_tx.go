package hermeskv

type StoreWithTx struct {
	store      StoreIface
	localState map[string]*ValMetaDeta
	isTxActive bool
}

var _ StoreIface = (*StoreWithTx)(nil)

func getStoreWithTx(capacity int) *StoreWithTx {
	return &StoreWithTx{
		// initialise the global state.
		store: GetNewKV(capacity),
	}
}

func (stx *StoreWithTx) Set(key string, value interface{}) error {
	// check if tx is active, if yes, route all ops to localState map
	if stx.isTxActive {
		stx.localState[key] = &ValMetaDeta{
			// for localState, value is the actual value being stored.
			value:       value,
			isTombStone: false,
		}
		return nil
	}

	return stx.store.Set(key, value)
}

func (stx *StoreWithTx) Get(key string) (interface{}, error) {
	if stx.isTxActive {
		valMeta, ok := stx.localState[key]

		if !ok {
			// if not found in localState then check globalState
			val, err := stx.store.Get(key)
			if err != nil {
				return "", err
			}

			return val, nil
		}

		// check if key is already marked as tombStone
		if valMeta.isTombStone {
			return "", ErrNoKey
		}

		return valMeta.value, nil
	}
	return stx.store.Get(key)
}

func (stx *StoreWithTx) Delete(key string) error {
	if stx.isTxActive {
		// first check whether the key is even present or not
		val, err := stx.Get(key)
		if err != nil {
			return ErrNoKey
		}

		// upsert the key with isTombStone as true in localState in case it isn't present
		// cause we need to record this op in the localState.
		stx.localState[key] = &ValMetaDeta{
			value:       val,
			isTombStone: true,
		}

		return nil
	}

	// this means Delete() is called outside out a tx. Yeet from the globalState directly.
	return stx.store.Delete(key)

}

func (stx *StoreWithTx) Begin() {
	// when begin is called, record all the operations of the ongoing tx
	// mark the current tx as active
	stx.isTxActive = true
	stx.localState = make(map[string]*ValMetaDeta)
}

func (stx *StoreWithTx) Commit() {
	// apply all DML ops to the globalState

	for k, v := range stx.localState {
		if v.isTombStone {
			// if this key is present in globalState then delete from it. otherwise ignore this op.
			if _, err := stx.store.Get(k); err == nil {
				stx.store.Delete(k)
			}

		} else {
			// apply normal operations
			stx.store.Set(k, v.value)
		}
	}
}

func (stx *StoreWithTx) Rollback() {
	// on rollback, clear the tx local temporary map holding all its operations
	clear(stx.localState)
}

func (stx *StoreWithTx) End() {
	// mark the ending of a tx
	stx.isTxActive = false
}
