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

//  查看某个peer节点已经安装实例化的链码
func systemByInstantiatedcc(c *g.Context) (list []string, err error) {

	log.Println("系统操作: 查看某个peer节点已经安装实例化的链码 => SystemByInstantiatedcc ")

	// 解析数据
	data, err := serializeBySystemByInstantiatedcc(c)
	if err != nil {
		log.Println("解析数据失败：", err)
	}

	log.Println("序列化成功：", data)

	// 查看某个peer节点已经安装实例化的链码
	arr, err1 := xnEnumerateExistingNodesByInsite(data)
	if err1 != nil {
		log.Println("查询节点已经安装的链码 err:", err1)
		return nil, err1
	}

	return arr, nil
}

//  向某个节点安装链码
func systemByInstallCCDepend(c *g.Context) (list []string, err error) {

	log.Println("系统操作: 向某个节点安装链码 => SystemByInstantiatedcc ")

	// 解析数据
	data, err := serializeBySystemByInstallCCDepend(c)
	if err != nil {
		log.Println("解析数据失败：", err)
	}

	log.Println("序列化成功：", data)

	// 查看某个peer节点已经安装实例化的链码
	arr, err1 := xnEnumerateExistingNodesByInstallCCDepend(data)
	if err1 != nil {
		log.Println("查询节点已经安装的链码 err:", err1)
		return nil, err1
	}

	return arr, nil
}
