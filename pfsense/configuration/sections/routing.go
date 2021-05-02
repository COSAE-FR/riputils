package sections

import (
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
)

type Route struct {
	Network     string `xml:"network"`
	Gateway     string `xml:"gateway"`
	Description string `xml:"descr"`
}

type Gateway struct {
	Interface      string                    `xml:"interface"`
	Gateway        string                    `xml:"gateway"`
	Name           string                    `xml:"name"`
	Weight         string                    `xml:"weight"`
	Protocol       string                    `xml:"ipprotocol"`
	Description    string                    `xml:"descr"`
	MonitorDisable helpers.TrueIfPresentBool `xml:"monitor_disable"`
	ActionDisable  helpers.TrueIfPresentBool `xml:"action_disable"`
}
