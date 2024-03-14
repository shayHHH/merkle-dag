package merkledag

const (
	FILE = iota
	DIR
)

type Node interface {
	Size() uint64 //返回节点大小
	Type() int    //返回节点类型
}

type File interface {
	Node

	Bytes() []byte //获取节点数据
}

type Dir interface {
	Node

	It() DirIterator //返回目录迭代器
}

type DirIterator interface {
	Next() bool //下一个节点

	Node() Node //返回当前节点功能
}
