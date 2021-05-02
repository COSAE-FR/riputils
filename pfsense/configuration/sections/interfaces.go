package sections

import (
	"encoding/xml"
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
	"io"
)

type Vlan struct {
	If          string                    `xml:"if"`
	Tag         uint16                    `xml:"tag"`
	PCP         helpers.TrueIfPresentBool `xml:"pcp"`
	Description string                    `xml:"descr"`
	VlanIf      string                    `xml:"vlanif"`
}

type IfGroup struct {
	Name        string                     `xml:"ifname"`
	Description string                     `xml:"descr"`
	Members     helpers.SpaceSeparatedList `xml:"members"`
}

type Interface struct {
	If          string                    `xml:"if"`
	Enable      helpers.TrueIfPresentBool `xml:"enable"`
	BlockBogons helpers.TrueIfPresentBool `xml:"blockbogons"`
	SpoofMac    helpers.TrueIfPresentBool `xml:"spoofmac"`
	Description string                    `xml:"descr"`
	Ip          string                    `xml:"ipaddr"`
	Subnet      uint8                     `xml:"subnet"`
	Gateway     string                    `xml:"gateway"`
}

type InterfaceList map[string]Interface

type XMLInterface struct {
	XMLName xml.Name
	Interface
}

func (l *InterfaceList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*l = InterfaceList{}
	for {
		var e XMLInterface

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*l)[e.XMLName.Local] = e.Interface
	}
	return nil
}

func (l InterfaceList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for name, iface := range l {
		i := XMLInterface{
			XMLName:   xml.Name{Local: name},
			Interface: iface,
		}
		err := e.Encode(i)
		if err != nil {
			return err
		}
	}
	return nil
}
