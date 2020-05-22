package model

// sdk
type Apper interface {
	Initialize()          // sdk
	CreateresMgmtClient() // msg
	CreateChannelCli()    // 通道
	CreateMspClient()     // msp
}

//  操作sdk
type ChannelDb interface {
	CreateChannel(path string, channelName string, org string, user string, orderers []string) error
	JoinChannel(path string, channelName string, org string, user string, orderers []string) error
}


//  peer 节点  

type PeerDb interface {


}