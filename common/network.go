package common

import (
	"errors"
	"net"
)

// GetIPForInterface provides rhe first configured IP of a network interface
func GetIPForInterface(interfaceName string) (ipAddress *net.IPNet, err error) {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if inter.Name == interfaceName {
			if addresses, err := inter.Addrs(); err == nil {
				for _, addr := range addresses {
					switch ip := addr.(type) {
					case *net.IPNet:
						if ip.IP.To4() != nil {
							return ip, nil
						}
					}
				}
			}
		}
	}
	return ipAddress, errors.New("no IP found")
}

func GetLocalIPStrings() (ips []string) {
	foo, err := net.InterfaceAddrs()

	if err == nil {
		for _, v := range foo {
			ips = append(ips, v.String())
		}
	}
	return
}

func GetLocalIPs() (ips []net.IP) {
	for _, ip := range GetLocalIPStrings() {
		localIP, _, err := net.ParseCIDR(ip)
		if err == nil {
			ips = append(ips, localIP)
		}
	}
	return
}

func GetLocalNetworks() (nets []net.IPNet) {
	for _, ip := range GetLocalIPStrings() {
		_, network, err := net.ParseCIDR(ip)
		if err == nil {
			nets = append(nets, *network)
		}
	}
	return
}