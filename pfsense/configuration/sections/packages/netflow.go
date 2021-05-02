package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type NetflowConfig struct {
	Enable           helpers.OnOffBool `xml:"enable"`
	CollectorAddress string            `xml:"collectoraddress"`
	CollectorPort    uint16            `xml:"collectorport"`
}

type NetflowCaptureConfig struct {
	Enable      helpers.OnOffBool `xml:"enable"`
	Interface   string            `xml:"interface"`
	Filter      string            `xml:"filter"`
	Description string            `xml:"description"`
}
