package controller

import (
	"github.com/gin-gonic/gin"
)

// 查询已经安装的链码
func QueryInstalledChaincode(c *gin.Context) {

	var s string
	res, err := queryInstalledChaincode(c)

	//  err
	if err != nil {
		s = "fail"
	} else {
		s = "succes"
	}
	GinBack(c, res, s)
	return
}

// Install// 查询已经安装实例化的链码
func QueryInstantiatedChaincode(c *gin.Context) {

	var s string
	res, err := queryInstantiatedChaincode(c)

	//  err
	if err != nil {
		s = "fail"
	} else {
		s = "succes"
	}
	GinBack(c, res, s)
	return
}
