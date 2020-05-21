package api

import "ki-sdk/util/action"

/**
@@@ 四级函数定义文件 ：暴露可以用的api 接口给 外部应用
*/

/**
@ SelectEncrypt 对数据进行加密
@ 2019年10月18日11:34:03
@ lidongsun
@ {参数：加密标识,加密数据（string）返回：hash}
*/
func GetcurrentHash(edType string, text string) (hash string) {
	return action.SelectEncrypt(edType, text)
}

//
func GetSliceArgs(funcName string, dataType string, storeData string, upDate string) (list []string, err error) {
	return action.SliceArgs(funcName, dataType, storeData, upDate)
}

// 参数分隔
func Cargsplict(args []string) (arg1 string, args2 [][]byte) {
	return action.Cargsplict(args)
}

// 查询参数
func GetSliceArgsById(funcName string, uuid string) (list []string, err error) {
	return action.SliceArgsById(funcName, uuid)
}

// 获取 服务器 Mem 使用率
func GetMemSize() (rep string) {
	return action.GetMemSize()
}

//	 公共方法 -- 获取 CPU 使用情况
func GetCpuSize() (rep string) {
	return action.GetCpuSize()
}

// -- 数据处理
func ArgsSplicing(arg []string) (res [][]byte, err error) {
	return action.ArgsSplicing(arg)
}
