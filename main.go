package main

import (

	m "ki-sdk/model"
	"ki-sdk/e2e"
	"ki-sdk/configless"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {

	// App
	m.InitSDK()
 
		e2e.SetupAndRuning( nil,
			fabsdk.WithEndpointConfig(configless.endpointConfigImpls...),
			fabsdk.WithCryptoSuiteConfig(configless.cryptoConfigImpls...),
			fabsdk.WithIdentityConfig(configless.identityConfigImpls...),
			fabsdk.WithMetricsConfig(configless.operationsConfigImpls...),
		)
	}

	// // 初始化路由
	// egg := r.InitRouter()

	// // 启动  server
	// err := egg.Run(":8080")
	// if err == nil {
	// 	log.Println("egg is starting")
	// } else {
	// 	log.Println("egg is err:", err)
	// }
}
