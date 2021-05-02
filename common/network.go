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
