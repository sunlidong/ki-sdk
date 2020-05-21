package model

import (
	"log"
)

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
		ConfigFile:      Init_ConfigFile,
		ChannelID:       Init_ChannelID,
		ChannelConfig:   Init_ChannelConfig,
		OrgAdmin:        Init_OrgAdmin,
		ChainCodeID:     Init_ChainCodeID,
		OrgName:         Init_OrgName,
		UserName:        Init_UserName,
		Version:         Init_Version,
		ChaincodeGoPath: Init_ChaincodeGoPath,
		GoPath:          Init_GoPath,
		OrdererID:       Init_OrdererID,
		Args:            Init_Args,
	}

	App = Application{
		SDK: &fSetUp,
	}
	return &fSetUp
}
