package action

var App Application

//  系统函数

/*
/***********************************************************************************************
*函数名 ： Fabric_Init
*函数功能描述 ：SDK 数据业务初始化 || 系统 数据
*函数参数 ：无
*函数返回值 ：
*作者 ：孙利栋
*函数创建日期 ：2019年10月14日13:53:37
*函数修改日期 ：
*修改人 ：
*修改原因 ：
*版本 ：V01
*历史版本 ：V01
***********************************************************************************************/
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

//func ccheckUser(hash string) (*msp.IdentityResponse, error){
//return  checkUser(hash)
//}
func cpeer() []string {
	//
	return peer()
}

// 调用 获取网络中所有节点
func getConfigureNodesTheNetwork() (peerNum int) {
	return ConfigureNodesTheNetwork()
}

func getConfigureOrgTheNetwork() (OrgNum int) {
	return ConfigureOrgTheNetwork()
}
