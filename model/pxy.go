package model

import (
	"log"
)

var App Application

/**
@ 对外 函数 || SDK 单个 实例化
*/
func InitSDK() error {
	var err error

	//	01.	sdk
	SDK := initsdk()

	//	02.	init
	err = SDK.Initialize()
	if err != nil {
		return err
	}
	log.Println("01 || 客戶端初始化成功")
	//	03.	msg
	err = SDK.CreateresMgmtClient()
	if err != nil {
		return err
	}
	log.Println("02 || 资源客戶端初始化成功")

	//	04.	cli
	err = SDK.CreateChannelCli()
	if err != nil {
		return err
	}
	log.Println("03 || 通道客戶端初始化成功")

	//	05.	msp
	err = SDK.CreateMspClient()
	if err != nil {
		return err
	}
	log.Println("04 || 证书客戶端初始化成功")

	return nil
}

func initsdk() *SDK {
	// 01. 声明sdk结构体对象
	fSetUp := SDK{
		ConfigFile:  ConfigFile,
		ChannelID:   ChannelID,
		OrgAdmin:    OrgName,
		ChainCodeID: ChainCodeID,
		OrgName:     OrgAdmin,
		UserName:    OrgAdmin,
		Version:     Version,
		OrdererID:   OrdererID,
	}

	App = Application{
		SDK: &fSetUp,
	}
	return &fSetUp
}
