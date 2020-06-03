package controller

import (
	m "ki-sdk/model"
)

/*
	------------- 调用  model 层 操作
*/

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
func createChannelConnection(p *SystemByJoinChannelDb) error {

	err := m.JoinChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)

	return err
}
