package router

import (
	c "ki-sdk/controller"

	g "github.com/gin-gonic/gin"
)

/*
	路由配置设置：
				1. 数据上链 	：	set
				2. 数据查询 	:	query
				3. 组织管理		:	org
				4. 配置查询		:	conf
				5. CA管理		:	ca
				6. 区块链浏览器	：	web
				6. 区块链浏览器	：	sql

*/

func InitRouter() *g.Engine {

	router := g.Default()
	{
		//路由组
		api := router.Group("/api")
		//路由组 v1  上链相关接口
		v1 := api.Group("/v1")

		wei := v1.Group("/wei")
		cha := v1.Group("/channel")

		// 数据上链
		{
			// 上链
			wei.POST("/up", c.UpLoad)

			//  查询
			wei.POST("/query", c.Load)
		}

		// channel -----------------------------  列表
		{
			cha.POST("/queryInstalledChaincode", c.QueryInstalledChaincode)
		}

		return router
	}
}
