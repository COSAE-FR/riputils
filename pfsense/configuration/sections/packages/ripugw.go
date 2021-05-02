package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type RipUgwConfig struct {
	Enable        helpers.OnOffBool `xml:"enable"`
	UnifiUrl      string            `xml:"unifi_url"`
	WanInterface  string            `xml:"wan_interface"`
	LanInterface  string            `xml:"lan_interface"`
	Wan2Interface string            `xml:"wan2_interface"`
	UidInterface  string            `xml:"uid_interface"`
	LogLevel      string            `xml:"log_level"`
}
