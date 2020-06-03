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
