package practitioners

// sdk

type practitioners interface {
	new()
	withEndpointConfig()
	withCryptoSuiteConfig()
	withIdentityConfig()
	withMetricsConfig()
	getName()
}

// e2e.SetupAndRuning(false, nil,
// 	fabsdk.WithEndpointConfig(endpointConfigImpls...),
// 	fabsdk.WithCryptoSuiteConfig(cryptoConfigImpls...),
// 	fabsdk.WithIdentityConfig(identityConfigImpls...),
// 	fabsdk.WithMetricsConfig(operationsConfigImpls...),
// )

//  生成4个参数

type practitionersByAtomicity interface {
	newWithEndpointConfig()
	newwithCryptoSuiteConfig()
	newwithIdentityConfig()
	newwithMetricsConfig()
	InitApi()
	initApi()
	setChannelsConfig()
	setOrgsConfig()
	setOrderersConfig()
	setpeersConfig()
	peersByLocalURL()
	caConfigObj()
	networkConfig()
	setPeerName()
	setOrdererName()
	setChannelName()
}
