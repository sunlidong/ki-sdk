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

// cha xun yi jing an zhuang de lian ma
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
