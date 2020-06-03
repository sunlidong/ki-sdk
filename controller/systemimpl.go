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
		log.Println("创建通道操作", err)
	}

	// 创建通道
	err = createChannelConnection(data)
	return err
}

// jia ru tong dao
func systemByJoinChannel(c *g.Context) (err error) {

	log.Println("系统操作=>systemByJoinChannel")

	// 解析数据
	data, err := serializeBySystemByJoinChannel(c)
	if err != nil {
		log.Println("解析数据失败：", err)
	}

	log.Println("序列化成功：", data)

	if err != nil {
		log.Println("创建通道操作", err)
	}

	// 创建通道
	err = channelPropertyAccess(data)
	return err
}

//  cha xun jie dian yi jing an zhuang de lian ma
func systemByXnNodeInfoListFree(c *g.Context) (list []string, err error) {

	log.Println("系统操作: 查询节点已经安装的链码 => xnNodeInfoListFree ")

	// 解析数据
	data, err := serializeBySystemByXnNodeInfoListFree(c)
	if err != nil {
		log.Println("解析数据失败：", err)
	}

	log.Println("序列化成功：", data)

	if err != nil {
		log.Println("创建通道操作", err)
	}

	// 查询 某个节点 已经安装 的链码
	arr, err1 := xnEnumerateExistingNodes(data)
	if err1 != nil {
		log.Println("查询节点已经安装的链码 err:", err1)
		return nil, err1
	}

	return arr, nil
}
