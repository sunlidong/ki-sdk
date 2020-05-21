package action

/**
@@@ 二级函数定义文件 ：调用函数 对内
*/

// 调用 sha512 对字符串进行处理
func cencDsha512(text string) (textHash string) {
	return encDsha512(text)
}

// 调用 sha384 对字符串进行处理
func cencDsha384(text string) (textHash string) {
	return encDsha384(text)
}

// 调用 sha512_224 对字符串进行处理
func cecncDsha512_224(text string) (textHash string) {
	return encDsha512_224(text)
}

// 调用 sha256 对字符串进行处理
func cencDsha256(text string) (textHash string) {
	return encDsha256(text)
}

// 调用 sha256_224 对字符串进行处理
func cencDsha256_224(text string) (textHash string) {
	return encDsha256_224(text)
}

// 调用 sha1 对字符串进行处理
func cencDsha1(text string) (textHash string) {
	return encDsha1(text)
}

// 调用 md5 对字符串进行处理
func cencDmd5(text string) (textHash string) {
	return encDmd5(text)
}
