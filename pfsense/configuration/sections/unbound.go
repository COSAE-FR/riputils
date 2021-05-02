package sections

import (
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
)

type UnboundHost struct {
	Host        string                     `xml:"host"`
	Domain      string                     `xml:"domain"`
	Ip          string                     `xml:"ip"`
	Description string                     `xml:"descr"`
	Aliases     helpers.CommaSeparatedList `xml:"aliases"`
}
type Unbound struct {
	Enable        helpers.TrueIfPresentBool  `xml:"enable"`
	Interfaces    helpers.CommaSeparatedList `xml:"active_interface"`
	OutInterface  helpers.CommaSeparatedList `xml:"outgoing_interface"`
	CustomOptions string                     `xml:"custom_options"`
	HideIdentity  helpers.TrueIfPresentBool  `xml:"hideidentity"`
	HideVersion   helpers.TrueIfPresentBool  `xml:"hideversion"`
	Hosts         []UnboundHost              `xml:"hosts"`
	MsgCacheSize  uint16                     `xml:"msgcachesize"`
	LogVerbosity  uint8                      `xml:"log_verbosity"`
}
