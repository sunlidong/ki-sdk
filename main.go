package main

import (
	m "ki-sdk/model"
	r "ki-sdk/router"
	"ki-sdk/test/integration"
	"log"
)

func main() {

	// App
	m.InitSDK()

	//
	integration.E2EBlockInit()

	// 初始化路由
	egg := r.InitRouter()

	// 启动  server
	err := egg.Run(":8080")
	if err == nil {
		log.Println("egg is starting")
	} else {
		log.Println("egg is err:", err)
	}
}
