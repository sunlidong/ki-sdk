package controller

import (
	m "ki-sdk/model"
)

/*
	------------- 调用  model 层 操作
*/

func createChannelConnection(p *SystemByCreateChannelDb) error {

	//1 .调用  model 层  操作
	// func CreateChannel(path string, channelName string, org string, user string, orderers []string) error {

	err := m.CreateChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)
	return err
}
