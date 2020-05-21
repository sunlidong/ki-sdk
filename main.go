package main

/**
* Main
* @lidong sun
* @Time 2019年10月14日10:25:50
*
 */
import (
	r_action "ki-sdk/router/action"
	r_api "ki-sdk/router/api"
)

func main() {

	server := r_api.InitRouter()

	server.Run(r_action.Port_01)
	log.

		//	03. 提示启动成功
		log.Println("服务启动成功.....")
}
