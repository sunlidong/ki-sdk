package model

// sdk
type Apper interface {
	Initialize()          // sdk
	CreateresMgmtClient() // msg
	CreateChannelCli()    // 通道
	CreateMspClient()     // msp
}
