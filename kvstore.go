package merkledag

type KVStore interface {
	Has(key []byte) (bool, error)   //是否存在
	Put(key, value []byte) error    //存
	Get(key []byte) ([]byte, error) //获取
	Delete(key []byte) error        //删除
}
