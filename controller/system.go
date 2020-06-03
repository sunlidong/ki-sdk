package controller

import (
	g "github.com/gin-gonic/gin"
)

// 创建通道
func SystemByCreateChannel(c *g.Context) {

	err := systemByCreateChannel(c)

	//  err
	if err != nil {
		GinBack(c, err, "fail")
	} else {
		GinBack(c, err, "success")
	}
	return
}

// 加入通道
func SystemByJoinChannel(c *g.Context) {

	err := systemByJoinChannel(c)

	//  err
	if err != nil {
		GinBack(c, err, "fail")
	} else {
		GinBack(c, err, "success")
	}
	return
}

/*
 cha xun yi jing an zhuang de lian ma

 {
	"configFile":"/mnt/d/gopath/src/ki-sdk/config/org1_peer0_admin.yaml",
	"org":"Org2",
	"peerHost":"peer1.org2.bookstore.com"

}
*/
func SystemByXnNodeInfoListFree(c *g.Context) {

	arr, err := systemByXnNodeInfoListFree(c)

	//  err
	if err != nil {
		GinBack(c, arr, "fail")
	} else {
		GinBack(c, arr, "success")
	}
	return
}
