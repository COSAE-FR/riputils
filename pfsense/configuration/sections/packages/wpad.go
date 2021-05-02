package packages

import "github.com/COSAE-FR/riputils/pfsense/configuration/helpers"

type WpadConfig struct {
	Enable      helpers.TrueIfPresentBool `xml:"enable"`
	Interface   string                    `xml:"interface"`
	ListenPort  uint16                    `xml:"listenport"`
	Workers     uint8                     `xml:"workers"`
	Description string                    `xml:"description"`
	Sgerror     helpers.OnOffBool         `xml:"sgerror"`
	E2gerror    helpers.OnOffBool         `xml:"e2gerror"`
	FileData    helpers.Base64String      `xml:"filedata"`
}
