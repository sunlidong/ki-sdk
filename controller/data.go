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

// 查询 peer节点已经安装实例化的链码
type PeerInstalledChaincode struct {
	PeerName    []string `json:"peerName"`
	ChannelName string   `json:"channelName"`
}

type AutoGenerated struct {
	LocalhostRep    string               `json:"localhostRep"`
	DNSMatchRegX    string               `json:"dnsMatchRegX"`
	Client          Client               `json:"client"`
	ChannelsConfig  []ChannelsConfigList `json:"channelsConfigs"`
	OrgsConfig      []OrgsConfig         `json:"orgsConfigs"`
	OrderersConfig  []OrderersConfig     `json:"orderersConfigs"`
	PeersConfig     []PeersConfig        `json:"peersConfigs"`
	PeersByLocalURL []PeersByLocalURL    `json:"peersByLocalURLs"`
	CaConfigObj     []CaConfigObj        `json:"caConfigObjs"`
}

type Client struct {
	Organization    string   `json:"organization"`
	Logging         string   `json:"logging"`
	CryptoConfig    string   `json:"cryptoConfig"`
	CredentialStore string   `json:"credentialStore"`
	TLSCerts        TLSCerts `json:"tlsCerts"`
}

type TLSCerts struct {
	Key  string `json:"Key"`
	Cert string `json:"Cert"`
}

type ChannelsConfigList struct {
	ChannelName    string         `json:"channelName"`
	ChannelsConfig ChannelsConfig `json:"channelsConfig"`
}

type ChannelsConfig struct {
	Orderers []string `json:"orderers"`
	Peers    []string `json:"peers"`
	Policies Policies `json:"policies"`
}

type Policies struct {
	MinResponses int       `json:"minResponses"`
	MaxTargets   int       `json:"maxTargets"`
	RetryOpts    RetryOpts `json:"retryOpts"`
}

type RetryOpts struct {
	Attempts       int     `json:"attempts"`
	InitialBackoff int     `json:"initialBackoff"`
	MaxBackoff     int     `json:"maxBackoff"`
	BackoffFactor  float64 `json:"backoffFactor"`
}

type OrgsConfig struct {
	Org        string   `json:"org"`
	Type       string   `json:"type"`
	CryptoPath string   `json:"cryptoPath"`
	Peers      []string `json:"peers"`
}

type OrderersConfig struct {
	OrderName        string `json:"orderName"`
	URL              string `json:"url"`
	Ssl              string `json:"ssl"`
	KeepAliveTime    int    `json:"keep-alive-time"`
	KeepAliveTimeout int    `json:"keep-alive-timeout"`
	KeepAlivePermit  bool   `json:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure"`
	TLSCaCert        string `json:"tlsCaCert"`
}

type PeersConfig struct {
	PeerURL          string `json:"peerUrl"`
	URL              string `json:"url"`
	Ssl              string `json:"ssl"`
	KeepAliveTime    int    `json:"keep-alive-time"`
	KeepAliveTimeout int    `json:"keep-alive-timeout"`
	KeepAlivePermit  bool   `json:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure"`
	TLSCaCert        string `json:"tlsCaCert"`
}

type peersByLocalURL struct {
	PeerURL          string `json:"peerUrl"`
	URL              string `json:"url"`
	Ssl              string `json:"ssl"`
	KeepAliveTime    int    `json:"keep-alive-time"`
	KeepAliveTimeout int    `json:"keep-alive-timeout"`
	KeepAlivePermit  bool   `json:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure"`
	TLSCaCert        string `json:"tlsCaCert"`
}

type CaConfigObj struct {
	CaOrgName    string `json:"caOrgName"`
	ID           string `json:"id"`
	URL          string `json:"url"`
	Path         string `json:"path"`
	Key          string `json:"key"`
	Cert         string `json:"cert"`
	EnrollID     string `json:"enrollID"`
	EnrollSecret string `json:"enrollSecret"`
	CaName       string `json:"caName"`
}
