package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type SplunkForwarderConfig struct {
	Enable     helpers.OnOffBool `xml:"enable"`
	DsHost     string            `xml:"ds_host"`
	DsPort     uint16            `xml:"ds_port"`
	DsInterval uint64            `xml:"ds_interval"`
	TlsCert    string            `xml:"tls_cert"`
}
