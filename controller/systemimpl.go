package controller

import (
	"log"

	g "github.com/gin-gonic/gin"
)

/**
 * @description:
 * @param {type}
 * @return:
 */
func systemByCreateChannel(c *g.Context) (err error) {

	log.Println("创建通道操作=>systemByCreateChannel")

	// 解析数据
	data, err := serializeBySystemByCreateChannel(c)
	if err != nil {
		log.Println("解析数据失败：", err)
	}

	log.Println("序列化成功：", data)

	if err != nil {
		log.Println("创建通道操作", err1)
	}

	// 创建通道
	err = createChannelConnection(&data)
	return err
}
