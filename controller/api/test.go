package api

import (
	"ki-sdk/controller/action"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 数据上链
func UpLoadtest(c *gin.Context) {
	list, err := action.CupLoadtest(c)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

//
// 微分格数据上链  2019年11月7日18:55:44
func T_UpLoadByBox(c *gin.Context) {
	//
	// 数据上链
	list, err := action.UpLoadByBox(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

// 微分格 -- 注册
func T_AddOrg(c *gin.Context) {
	// 查询
	result, err := action.T_caddOrg(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   result,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

//
func T_SetUser(c *gin.Context) {
	// 查询
	result, err := action.T_setUser(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

// 测试链码
func TestUpload(c *gin.Context) {
	// 查询
	result, err := action.TestUpload(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

func TestSelect(c *gin.Context) {
	// 查询
	result, err := action.TestSelect(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}
