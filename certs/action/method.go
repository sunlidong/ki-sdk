package action

// 初始化实例
func initForCerts() (rep *Setup) {
	fsetup := Setup{
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

	return &fsetup
}
