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
