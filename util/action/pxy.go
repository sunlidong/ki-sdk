package action

/**
@@@ 三级函数定义文件 ：调用函数 对外  ，暴露给API
*/

/**
@ upeEncrypt  对 api 暴露 函数 对数据进行加密
@ 2019年10月18日11:04:13
@ No:1101
*/
func SelectEncrypt(index string, text string) (hash string) {
	//
	switch index {
	case EncryptText(EncryptDsha512):
		hash = cencDsha512(text)
	case EncryptText(EncryptDsha384):
		hash = cencDsha384(text)
	case EncryptText(EncryptDsha512_224):
		hash = cecncDsha512_224(text)
	case EncryptText(EncryptDsha256):
		hash = cencDsha256(text)
	case EncryptText(EncryptDsha256_224):
		hash = cencDsha256_224(text)
	case EncryptText(EncryptDsha1):
		hash = cencDsha1(text)
	case EncryptText(EncryptDmd5):
		hash = cencDmd5(text)
	default:
		hash = EncryptText(EncryptNull)
	}

	//返回

	return hash

}

//
func SliceArgs(funcName string, dataType string, storeData string, upDate string) (list []string, err error) {
	return sliceArgs(funcName, dataType, storeData, upDate)
}

//
func Cargsplict(args []string) (arg1 string, args2 [][]byte) {
	return argsplict(args)
}

//
func SliceArgsById(funcName string, uuid string) (list []string, err error) {
	return sliceArgsById(funcName, uuid)
}

// 获取 服务器 Mem 使用率
func GetMemSize() (rep string) {
	return getMemSize()
}

// 获取 服务器 CPU 使用率
func GetCpuSize() (rep string) {
	return getCpuSize()
}

// -- 数据处理
func ArgsSplicing(arg []string) (res [][]byte, err error) {
	return argsSplicing(arg)
}
