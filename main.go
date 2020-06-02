package main

import (
	"ki-sdk/install"
	m "ki-sdk/model"
	r "ki-sdk/router"
	"log"
)

func main() {

	// App
	m.InitSDK()
	install.InitApi()

	// App// 初始化路由
	egg := r.InitRouter()

	// 启动  server
	err := egg.Run(":8080")
	if err == nil {
		log.Println("egg is starting")
	} else {
		log.Println("egg is err:", err)
	}

}
