package controller

import (
	"fmt"
	"net/http"

	g "github.com/gin-gonic/gin"
)

// 数据上链
func UpLoad(c *g.Context) {

	fmt.Println("数据上链操作")
	c.JSON(
		http.StatusOK,
		g.H{"status": "200",
			"data": "test",
		})
}
