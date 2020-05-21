package model

// sdk
type App interface {
	Initialize()          // sdk
	CreateresMgmtClient() // msg
	CreateChannelCli()    // 通道
	CreateMspClient()     // msp
}
