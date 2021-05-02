/*
 * Copyright (c) 2021. Gaetan Crahay
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file or at
 * https://opensource.org/licenses/MIT.
 */

package configuration

import (
	"encoding/xml"
	"github.com/COSAE-FR/riputils/pfsense/configuration/sections"
	"github.com/COSAE-FR/riputils/pfsense/configuration/sections/packages"
)

type VersionConfiguration struct {
	XMLName     xml.Name `xml:"pfsense"`
	FileVersion string   `xml:"version"`
	Revision    Revision `xml:"revision"`
}

type BaseConfiguration struct {
	VersionConfiguration
	System      sections.System        `xml:"system"`
	Interfaces  sections.InterfaceList `xml:"interfaces"`
	Vlans       []sections.Vlan        `xml:"vlans>vlan"`
	Routes      []sections.Route       `xml:"staticroutes>route"`
	IfGroups    []sections.IfGroup     `xml:"ifgroups>ifgroupentry"`
	Unbound     sections.Unbound       `xml:"unbound"`
	Syslog      sections.Syslog        `xml:"syslog"`
	Gateways    []sections.Gateway     `xml:"gateways>gateway_item"`
	GatewayIpv4 string                 `xml:"gateways>defaultgw4"`
	GatewayIpv6 string                 `xml:"gateways>defaultgw6"`
	SysCtls     []sections.SysCtl      `xml:"sysctl>item"`
	Dhcp        sections.DhcpList      `xml:"dhcpd"`
}

type PackagesConfiguration struct {
	BaseConfiguration
	Packages packages.KnownPackagesConfig `xml:"installedpackages"`
}

type Revision struct {
	Time        string `xml:"time"`
	Description string `xml:"description"`
	Username    string `xml:"username"`
}
