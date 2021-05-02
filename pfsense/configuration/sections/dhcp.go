package sections

import (
	"encoding/xml"
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
	"io"
)

type Dhcp struct {
	Range                  DhcpRange                 `xml:"range"`
	FailoverPeerip         string                    `xml:"failover_peerip"`
	Defaultleasetime       string                    `xml:"defaultleasetime"`
	Maxleasetime           string                    `xml:"maxleasetime"`
	Netmask                string                    `xml:"netmask"`
	Gateway                string                    `xml:"gateway"`
	Domain                 string                    `xml:"domain"`
	Domainsearchlist       string                    `xml:"domainsearchlist"`
	Ddnsdomain             string                    `xml:"ddnsdomain"`
	Ddnsdomainprimary      string                    `xml:"ddnsdomainprimary"`
	Ddnsdomainkeyname      string                    `xml:"ddnsdomainkeyname"`
	Ddnsdomainkeyalgorithm string                    `xml:"ddnsdomainkeyalgorithm"`
	Ddnsdomainkey          string                    `xml:"ddnsdomainkey"`
	MacAllow               string                    `xml:"mac_allow"`
	MacDeny                string                    `xml:"mac_deny"`
	Ddnsclientupdates      string                    `xml:"ddnsclientupdates"`
	Tftp                   string                    `xml:"tftp"`
	Ldap                   string                    `xml:"ldap"`
	Nextserver             string                    `xml:"nextserver"`
	Filename               string                    `xml:"filename"`
	Filename32             string                    `xml:"filename32"`
	Filename64             string                    `xml:"filename64"`
	Rootpath               string                    `xml:"rootpath"`
	Enable                 helpers.TrueIfPresentBool `xml:"enable"`
	Dhcpleaseinlocaltime   string                    `xml:"dhcpleaseinlocaltime"`
	Ntpserver              string                    `xml:"ntpserver"`
	Staticarp              string                    `xml:"staticarp"`
	Denyunknown            string                    `xml:"denyunknown"`
	Numberoptions          []DhcpOption              `xml:"numberoptions>item,omitempty"`
}

type DhcpOption struct {
	Number string `xml:"number"`
	Type   string `xml:"type"`
	Value  string `xml:"value"`
}

type DhcpRange struct {
	From string `xml:"from"`
	To   string `xml:"to"`
}

type DhcpList map[string]Dhcp

type XMLDhcp struct {
	XMLName xml.Name
	Dhcp
}

func (l *DhcpList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*l = DhcpList{}
	for {
		var e XMLDhcp

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*l)[e.XMLName.Local] = e.Dhcp
	}
	return nil
}

func (l DhcpList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for name, dhcp := range l {
		i := XMLDhcp{
			XMLName: xml.Name{Local: name},
			Dhcp:    dhcp,
		}
		err := e.Encode(i)
		if err != nil {
			return err
		}
	}
	return nil
}
