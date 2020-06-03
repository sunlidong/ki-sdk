package controller

import (
	m "ki-sdk/model"
)

/*
	------------- 调用  model 层 操作
*/

// chuang jian tong dao
func createChannelConnection(p *SystemByCreateChannelDb) error {

	err := m.CreateChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)

	return err
}

// jia  ru  tong  dao
func channelPropertyAccess(p *SystemByJoinChannelDb) error {

	err := m.JoinChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)

	return err
}


// type SystemByXnNodeInfoListFreeDb struct {
// 	PeerHost   string `json:"peerHost"`
// 	ConfigFile string `json:"configFile"`
// 	Org        string `json:"org"`
// }

// 查询 某个 节点 已经 实例化 的 链码
func xnEnumerateExistingNodes(p *SystemByXnNodeInfoListFreeDb) (list []string, err error ){

return	func(func()m.ResmgmtClient{
			return &m.ResmgmtClient{
			Client:func()*resmgmt.Client{
				resmgmt_Client,err:= m.GetResmgmtClient(
					p.ConfigFile,
					p.Org
				 )
				 if err !=nil{
					log.Println("实例化 resmgmt_Client 失败： ",err)
		
				 }
				return resmgmt_Client
			}()
		}
	}())([]string,error){
		// TODO
		arr,err:=d.GetInstalledCC(p.PeerHost)
		if err != nil {
			return nil, err
		}
		return  arr ,nil
	}()

}
