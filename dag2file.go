package merkledag

// Hash to file
//从KVStore中读取与给定哈希值对应的数据，然后根据路径返回相应的文件内容
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 根据hash和path， 返回对应的文件, hash对应的类型是tree

	// 通过KVStore获取与给定哈希值对应的数据
	data := store.Get(hash)

	// 解析路径，获取文件夹和文件名
	dirPath, fileName := parsePath(path)

	// 根据文件夹路径获取Dir对象
	dir := store.GetDir(dirPath)

	// 遍历文件夹，查找对应的文件
	for it := dir.It(); it.HasNext(); {
		node := it.Next()

		// 判断是否为File类型且文件名匹配
		if node.Type == File && node.Name == fileName {
			// 返回文件内容
			return data
		}
	}
	return nil
}
