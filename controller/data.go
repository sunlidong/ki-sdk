package controller

//
type ChainDb struct {
	ChannelName   string   `json:"channelName"`
	ChainCodeName string   `json:"chainCodeName"`
	FunctionName  string   `json:"functionName"`
	Data          []string `json:"data"`
}
