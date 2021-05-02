package packages

import (
	"encoding/xml"
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
	"io"
)

type PackageTab struct {
	Text   string `xml:"text"`
	URL    string `xml:"url"`
	Active string `xml:"active"`
}

type Package struct {
	Name               string       `xml:"name"`
	Description        string       `xml:"descr"`
	PkgInfolink        string       `xml:"pkginfolink"`
	Configurationfile  string       `xml:"configurationfile"`
	Tabs               []PackageTab `xml:"tabs>tab"`
	IncludeFile        string       `xml:"include_file"`
	InternalName       string       `xml:"internal_name"`
	Website            string       `xml:"website"`
	FilterRuleFunction string       `xml:"filter_rule_function"`
	AfterInstallInfo   string       `xml:"after_install_info"`
}

type Service struct {
	Name           string                    `xml:"name"`
	Description    string                    `xml:"description"`
	Executable     string                    `xml:"executable"`
	RcFile         string                    `xml:"rcfile"`
	StatusCommand  string                    `xml:"custom_php_service_status_command,omitempty"`
	RestartCommand string                    `xml:"restartcmd,omitempty"`
	StopCommand    string                    `xml:"stopcmd"`
	StartCommand   string                    `xml:"startcmd"`
	StartsOnSync   helpers.TrueIfPresentBool `xml:"starts_on_sync"`
}

type Menu struct {
	Name       string `xml:"name"`
	Tooltip    string `xml:"tooltiptext"`
	Section    string `xml:"section"`
	ConfigFile string `xml:"configfile"`
	URL        string `xml:"url"`
}

type ConfigsMap map[string][]map[string]string

func (l *ConfigsMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *l == nil {
		*l = make(map[string][]map[string]string)
	}
	(*l)[start.Name.Local] = []map[string]string{}
	for {
		var e XMLConfig
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		values := make(map[string]string)
		for _, v := range e.Value {
			values[v.XMLName.Local] = v.Value
		}
		(*l)[start.Name.Local] = append((*l)[start.Name.Local], values)
	}
	return nil
}

func (l ConfigsMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// TODO
	return nil
}

type XMLConfig struct {
	XMLName xml.Name
	Value   []XMLConfigItem `xml:",any"`
}

type XMLConfigItem struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type BasePackageConfig struct {
	Details              []Package                    `xml:"package"`
	Services             []Service                    `xml:"service"`
	Menus                []Menu                       `xml:"menu"`
}

type KnownPackagesConfig struct {
	BasePackageConfig
	SplunkForwarder      SplunkForwarderConfig        `xml:"splunkforwarder>config"`
	Riprovision          RiprovisionConfig            `xml:"riprovision>config"`
	RiprovisionAuth      []RiprovisionAuthConfig      `xml:"riprovisionauth>config"`
	RiprovisionUsers     []RiprovisionUsersConfig     `xml:"riprovisionusers>config"`
	RiprovisionTemplates []RiprovisionTemplatesConfig `xml:"riprovisiontemplates>config"`
	RiprovisionDevices   []RiprovisionDevicesConfig   `xml:"riprovisiondevices>config"`
	RiprovisionMac       []string                     `xml:"riprovisionmac>config>mac"`
	RipUgw               RipUgwConfig                 `xml:"ripugw>config"`
	Netflow              NetflowConfig                `xml:"netflow>config"`
	NetflowCapture       []NetflowConfig              `xml:"netflowcaptures>config"`
	Riproxy              *RiproxyConfig               `xml:"riproxy>config"`
	RiproxyHttp          []RiproxyHttpConfig          `xml:"riproxyhttp>config"`
	RiproxyProxy         []RiproxyProxyConfig         `xml:"riproxyproxy>config"`
	Squid                SquidConfig                  `xml:"squid>config"`
	SquidCache           SquidCacheConfig             `xml:"squidcache>config"`
	SquidReverseGeneral  SquidReverseGeneralConfig    `xml:"squidreversegeneral>config"`
	SquidReversePeers    []SquidReversePeerConfig     `xml:"squidreversepeer>config"`
	SquidReverseUris     []SquidReverseUriConfig      `xml:"squidreverseuris>config"`
	Wpad                 []WpadConfig                 `xml:"wpad>config"`
	Configs              ConfigsMap                   `xml:",any"`
}
