package controller

type ChainDb struct {
	ChannelName   string   `json:"channelName"`
	ChainCodeName string   `json:"chainCodeName"`
	FunctionName  string   `json:"functionName"`
	Data          []string `json:"data"`
}

// 查询 peer节点已经安装的链码
type PeerInstallChaincode struct {
	PeerName []string `json:"peerName"`
}
