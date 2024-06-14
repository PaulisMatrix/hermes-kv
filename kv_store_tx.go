package hermeskv

type StoreTx struct {
	store      *Store
	isTxActive bool
}

func getStoreWithTx(capacity int) *StoreTx {
	return &StoreTx{
		store:      GetNewKV(capacity),
		isTxActive: false,
	}
}

func (stx *StoreTx) Begin() {

}

func (stx *StoreTx) Commit() {

}

func (stx *StoreTx) Rollback() {

}
