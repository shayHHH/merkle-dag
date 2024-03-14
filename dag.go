package merkledag

import "hash"

func Add(store KVStore, node Node, h hash.Hash) []byte {
	//将Node中的数据保存在KVStore中，并计算出Merkle Root
	// 根据节点类型不同进行处理
	switch node.Type() {
	case FILE:
		fileNode := node.(File)              // 类型断言以访问 File 接口特定的方法
		data := fileNode.Bytes()             // 从 File 节点获取字节数据
		store.Put([]byte("file_data"), data) // Save the file data in the KVStore
	case DIR:
		dirNode := node.(Dir) // 类型断言以访问 Dir 接口特定的方法
		// 迭代目录中的节点
		dirIterator := dirNode.It()
		for dirIterator.Next() {
			subNode := dirIterator.Node()
			// 递归调用 Add 函数处理每个子节点
			Add(store, subNode, h)
		}
	}

	// 根据保存在 KVStore 中的数据计算并返回 Merkle Root
	return calculateMerkleRoot(store, h)
}

func calculateMerkleRoot(store KVStore, h hash.Hash) []byte {
	// Get all keys and values from the KVStore
	keys, _ := store.GetAllKeys()
	var dataList [][]byte
	for _, key := range keys {
		value, _ := store.Get(key)
		dataList = append(dataList, value)
	}

	// Calculate the Merkle Root based on the stored data
	for _, data := range dataList {
		h.Write(data)
	}
	return h.Sum(nil)
}
