package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type RiprovisionConfig struct {
	Enable                 string                     `xml:"enable"`
	ListeningInterface     helpers.CommaSeparatedList `xml:"listening_interface"`
	LogLevel               string                     `xml:"log_level"`
	MaxDevices             uint64                     `xml:"max_devices"`
	ProvisioningInterfaces helpers.CommaSeparatedList `xml:"provisioning_interfaces"`
	ProvisioningSyslogPort uint16                     `xml:"provisioning_syslog_port"`
	DhcpEnable             helpers.OnOffBool          `xml:"dhcp_enable"`
	DhcpBaseNetwork        string                     `xml:"dhcp_base_network"`
	DhcpClientPrefix       uint8                      `xml:"dhcp_client_prefix"`
}

type RiprovisionAuthConfig struct {
	Type        string `xml:"type"`
	Password    string `xml:"password"`
	Description string `xml:"description"`
	Keyfile     string `xml:"keyfile"`
}

type RiprovisionUsersConfig struct {
	Description string `xml:"description"`
	Username    string `xml:"username"`
}
type RiprovisionTemplatesConfig struct {
	Code     string `xml:"code"`
	Template string `xml:"template"`
}
type RiprovisionDevicesConfig struct {
	Code     string `xml:"code"`
	Template string `xml:"template"`
}
