package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type RiproxyBaseWpadConfig struct {
	DirectInterfaces helpers.CommaSeparatedList `xml:"directinterfaces"`
	InterfaceDirect  helpers.OnOffBool          `xml:"listeningdirect"`
}

type RiproxyBaseProxyConfig struct {
	ProxyPort            uint16            `xml:"proxyport"`
	AllowHighPorts       helpers.OnOffBool `xml:"allowhigh"`
	AllowLowPorts        helpers.OnOffBool `xml:"allowlow"`
	BlockIps             helpers.OnOffBool `xml:"blockips"`
	BlockLocalServices   helpers.OnOffBool `xml:"blocklocal"`
	HttpTransparent      helpers.OnOffBool `xml:"httptransparent"`
	HttpsTransparentPort uint16            `xml:"httpstransparentport"`
	BlockByIdn           helpers.OnOffBool `xml:"blockbyidn"`
	Block                []string          `xml:"row>host"`
}

type RiproxyConfig struct {
	Enable   helpers.OnOffBool `xml:"enable"`
	LogLevel string            `xml:"loglevel"`
	RiproxyBaseWpadConfig
	RiproxyBaseProxyConfig
}

type RiproxyReverseProxyConfig struct {
	Host      string `xml:"host"`
	PeerIP    string `xml:"peerip"`
	PeerPort  uint16 `xml:"peerport"`
	Interface string `xml:"interface"`
}

type RiproxyHttpConfig struct {
	Interface            string            `xml:"interface"`
	Enable               helpers.OnOffBool `xml:"enable"`
	ExternalProxy        helpers.OnOffBool `xml:"externalproxy"`
	ExternalProxyAddress string            `xml:"proxyaddress"`
	RiproxyBaseWpadConfig
	ReverseProxies []RiproxyReverseProxyConfig `xml:"row"`
}

type RiproxyProxyConfig struct {
	Interface string            `xml:"interface"`
	Enable    helpers.OnOffBool `xml:"enable"`
	RiproxyBaseProxyConfig
}
