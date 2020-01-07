package action

//	实例化
func initSDK() {
	rep := InitForCerts()
	rep.Initialize()
	rep.CreateMsgClient()
	rep.CreateChannelCli()
	rep.CreateMspClient()

}
