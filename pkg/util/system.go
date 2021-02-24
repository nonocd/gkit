package util

import (
	"net"
	"os"
)

// GetHostName 获取主机名称
func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}

// GetIPv4 获取系统IP地址
func GetIPv4() (net.IP, bool) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, false
	}

	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, true
		}
	}
	return nil, false
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}
