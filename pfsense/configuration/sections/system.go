package sections

import (
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
)

type System struct {
	Hostname                      string                     `xml:"hostname"`
	Timezone                      string                     `xml:"timezone"`
	Domain                        string                     `xml:"domain"`
	DnsServers                    []string                   `xml:"dnsserver"`
	DnsAllowOverride              helpers.OnOffBool          `xml:"dnsallowoverride"`
	Timeservers                   helpers.CommaSeparatedList `xml:"timeservers"`
	NextUid                       uint32                     `xml:"nextuid"`
	NextGid                       uint32                     `xml:"nextgid"`
	Disablenatreflection          helpers.YesNoBool          `xml:"disablenatreflection"`
	Disablesegmentationoffloading helpers.TrueIfPresentBool  `xml:"disablesegmentationoffloading"`
	Disablelargereceiveoffloading helpers.TrueIfPresentBool  `xml:"disablelargereceiveoffloading"`
	Ipv6dontcreatelocaldns        helpers.TrueIfPresentBool  `xml:"ipv6dontcreatelocaldns"`
	PreferIpv4                    helpers.TrueIfPresentBool  `xml:"prefer_ipv4"`
	Maximumtableentries           uint64                     `xml:"maximumtableentries"`
	PowerdAcMode                  string                     `xml:"powerd_ac_mode"`
	PowerdBatteryMode             string                     `xml:"powerd_battery_mode"`
	PowerdNormalMode              string                     `xml:"powerd_normal_mode"`
	AlreadyRunConfigUpgrade       helpers.TrueIfPresentBool  `xml:"already_run_config_upgrade"`
	Language                      string                     `xml:"language"`
	Users                         []SystemUser               `xml:"user"`
	Groups                        []SystemGroup              `xml:"group"`
	ACB                           ACB                        `xml:"acb"`
	WebGui                        struct {
		Protocol           string                     `xml:"protocol"`
		SslCertref         string                     `xml:"ssl-certref"`
		Althostnames       helpers.SpaceSeparatedList `xml:"althostnames"`
		Dashboardcolumns   uint8                      `xml:"dashboardcolumns"`
		Authmode           string                     `xml:"authmode"`
		Api                helpers.TrueIfPresentBool  `xml:"api"`
		Interfaces         helpers.CommaSeparatedList `xml:"interfaces"`
		Webguicss          string                     `xml:"webguicss"`
		Logincss           string                     `xml:"logincss"`
		Webguihostnamemenu string                     `xml:"webguihostnamemenu"`
	} `xml:"webgui"`
}

type SystemUser struct {
	Scope                          string                    `xml:"scope"`
	BcryptHash                     string                    `xml:"bcrypt-hash"`
	Description                    string                    `xml:"descr"`
	Name                           string                    `xml:"name"`
	Expires                        string                    `xml:"expires"`
	DashboardColumns               uint8                     `xml:"dashboardcolumns"`
	AuthorizedKeys                 string                    `xml:"authorizedkeys"`
	IpsecPsk                       string                    `xml:"ipsecpsk"`
	Disabled                       helpers.TrueIfPresentBool `xml:"disabled"`
	Customsettings                 string                    `xml:"customsettings"`
	Webguicss                      string                    `xml:"webguicss"`
	Webguifixedmenu                string                    `xml:"webguifixedmenu"`
	Webguihostnamemenu             string                    `xml:"webguihostnamemenu"`
	Interfacessort                 string                    `xml:"interfacessort"`
	Dashboardavailablewidgetspanel string                    `xml:"dashboardavailablewidgetspanel"`
	Systemlogsfilterpanel          string                    `xml:"systemlogsfilterpanel"`
	Systemlogsmanagelogpanel       string                    `xml:"systemlogsmanagelogpanel"`
	Statusmonitoringsettingspanel  string                    `xml:"statusmonitoringsettingspanel"`
	Webguileftcolumnhyper          string                    `xml:"webguileftcolumnhyper"`
	Disablealiaspopupdetail        string                    `xml:"disablealiaspopupdetail"`
	Pagenamefirst                  string                    `xml:"pagenamefirst"`
	Uid                            uint32                    `xml:"uid"`
}

type SystemGroup struct {
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Scope       string   `xml:"scope"`
	Gid         uint32   `xml:"gid"`
	Member      []uint32 `xml:"member"`
	Priv        string   `xml:"priv"`
}

type Syslog struct {
	FilterDescriptions string                    `xml:"filterdescriptions"`
	Nentries           string                    `xml:"nentries"`
	Remoteserver       string                    `xml:"remoteserver"`
	Remoteserver2      string                    `xml:"remoteserver2"`
	Remoteserver3      string                    `xml:"remoteserver3"`
	SourceIp           string                    `xml:"sourceip"`
	Protocol           string                    `xml:"ipproto"`
	LogAll             helpers.TrueIfPresentBool `xml:"logall"`
	Enable             helpers.TrueIfPresentBool `xml:"enable"`
}

type SysCtl struct {
	Name        string `xml:"tunable"`
	Value       string `xml:"value"`
	Description string `xml:"descr"`
}

type ACB struct {
	Enable   helpers.YesNoBool `xml:"enable"`
	Hint     string            `xml:"hint"`
	Server   string            `xml:"server"`
	Password string            `xml:"encryption_password"`
}
