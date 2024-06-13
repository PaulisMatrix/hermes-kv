package hermeskv

type StoreTx struct {
	store *Store
}

func getStoreWithTx(capacity int) *StoreTx {
	return &StoreTx{
		store: GetNewKV(capacity),
	}
}

func (stx *StoreTx) Begin() {

}

func (stx *StoreTx) Commit() {

}

func (stx *StoreTx) Rollback() {

}
