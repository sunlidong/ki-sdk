package controller

import (
	g "github.com/gin-gonic/gin"
)

func serializeBySystemByCreateChannel(c *g.Context) (data *SystemByCreateChannelDb, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}

// xuliehua jiaru tongdao
func serializeBySystemByJoinChannel(c *g.Context) (data *SystemByJoinChannelDb, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}

// xuliehua jiaru tongdao
func serializeBySystemByXnNodeInfoListFree(c *g.Context) (data *SystemByXnNodeInfoListFreeDb, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}

//	查询某个节点已经实例化的链码
func serializeBySystemByInstantiatedcc(c *g.Context) (data *SystemByInstantiatedccDb, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}

//	向某个节点安装链码
func serializeBySystemByInstallCCDepend(c *g.Context) (data *SystemByInstallCCDependDb, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}
