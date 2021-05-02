package configuration

import (
	"errors"
	"github.com/COSAE-FR/riputils/pfsense/configuration/sections"
)

type InterfaceGetter interface {
	GetInterface(pfSenseInterface string) *sections.Interface
	GetPhysicalInterfaceName(pfSenseInterface string) (string, error)
}

func (c BaseConfiguration) GetInterface(pfSenseInterface string) *sections.Interface {
	config, ok := c.Interfaces[pfSenseInterface]
	if ok {
		return &config
	}
	return nil
}

func (c BaseConfiguration) GetPhysicalInterfaceName(pfSenseInterface string) (string, error) {
	config := c.GetInterface(pfSenseInterface)
	if config == nil {
		return "", errors.New("not found")
	}
	if len(config.If) == 0 {
		return "", errors.New("no physical interface")
	}
	return config.If, nil
}