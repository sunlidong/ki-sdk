package controller

import (
	"github.com/gin-gonic/gin"
)

// 数据上链
func QueryInstalledChaincode(c *gin.Context) {
	//
	// 数据上链
	var s string
	res, err := queryInstalledChaincode(c)

	//  err
	if err != nil {
		s = "fail"
	} else {
		s = "succes"
	}
	GinBack(res, s)
	return
}
