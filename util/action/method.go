package action

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"log"
)

/**
@@@ 一级函数定义文件 ：定义 函数本体
*/

/**
@ 加解密中心
@ 2019年10月18日10:29:17
@ lidongsun
*/

/**
@ Sha512
@ 2019年10月18日10:55:10
@ NO 1031
*/
func encDsha512(text string) (textHash string) {
	h := sha512.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
@ Sha512-384
@ 2019年10月18日10:55:17
@ NO 1032
*/
func encDsha384(text string) (textHash string) {
	h := sha512.New384()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
@ Sha512_224
@ 2019年10月18日10:35:49
@ NO 1033
*/
func encDsha512_224(text string) (textHash string) {
	h := sha512.New512_224()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
@ encDsha256
@ 2019年10月18日10:35:55
@ NO 1034
*/
func encDsha256(text string) (textHash string) {
	h := sha256.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
@ encDsha256_224
@ 2019年10月18日10:35:55
@ NO 1035
*/
func encDsha256_224(text string) (textHash string) {
	h := sha256.New224()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
@ encDsha1
@ 2019年10月18日10:35:55
@ NO 1036
*/
func encDsha1(text string) (textHash string) {
	h := sha1.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
@ encDmd5
@ 2019年10月18日10:35:55
@ NO 1037
*/
func encDmd5(text string) (textHash string) {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// 生成切片
func sliceArgs(funcName string, dataType string, storeData string, upDate string) (list []string, err error) {
	//以下参数是请求链码的相关配置

	var args []string
	if funcName == "" || dataType == "" || storeData == "" || upDate == "" {
		//
		return args, errors.New("args is nil ")
	}
	args = append(args, funcName)  // 函数名称
	args = append(args, dataType)  //数据类型
	args = append(args, storeData) //数据Data
	args = append(args, upDate)    //上链时间

	//返回
	return args, nil
}

// 生成切片 ||请求
func sliceArgsById(funcName string, uuid string) (list []string, err error) {
	//以下参数是请求链码的相关配置

	var args []string
	if uuid == "" || funcName == "" {
		//
		log.Println("uuid=>", uuid)
		log.Println("funcName=>", funcName)
		return args, errors.New("args is nil ")
	}
	args = append(args, funcName) // 函数名称
	args = append(args, uuid)     //uuid
	//返回
	return args, nil
}

/**
@ splic
@ 2019年10月18日17:32:44
@
*/
func argsplict(args []string) (arg1 string, args2 [][]byte) {
	//
	funcName := args[0]
	argsbyte := [][]byte{}
	//
	for i := 1; i < len(args); i++ {
		//
		argsbyte = append(argsbyte, []byte(args[i]))
	}
	return funcName, argsbyte
}

// -- 数据拼接
func argsSplicing(arg []string) (res [][]byte, err error) {
	//
	if len(arg) > 0 {
		for k, _ := range arg {
			res = append(res, []byte(arg[k]))
		}
		return res, nil
	} else {
		return nil, errors.New("arg is <= 0")
	}
}
