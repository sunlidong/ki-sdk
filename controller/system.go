package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	g "github.com/gin-gonic/gin"
)

// 创建通道
func SystemByCreateChannel(c *g.Context) {

	err := systemByCreateChannel(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": "fail",
				"data":   err,
			})
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": "success",
				"data": err,
			})
	}
	return
}
