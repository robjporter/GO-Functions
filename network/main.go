package network

import (
	"net"
	"strconv"
)

var (
	classAPrivate    = mustParseCIDR("10.0.0.0/8")
	classBPrivate    = mustParseCIDR("172.16.0.0/12")
	classCPrivate    = mustParseCIDR("192.168.0.0/16")
	ipv6UniqueLocal  = mustParseCIDR("fc00::/7")
	LoopbackIPv4CIDR = "127.0.0.0/8"
	LoopbackIPv6CIDR = "::1/128"
)

func IsIPAddress(ip string) bool {
	test := net.ParseIP(ip)
	if test.To4() == nil {
		return false
	} else {
		return true
	}
}

func mustParseCIDR(s string) *net.IPNet {
	_, net, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}
	return net
}

func GetMainIP() []string {
	ip := []string{}
	list, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range list {
		if iface.Name != "lo0" && iface.Name != "awdl0" && iface.Name != "utun0" && iface.Name != "utun1" &&
			iface.Name != "gif0" && iface.Name != "stf0" && iface.Name != "bridge0" && iface.Name != "p2p0" {
			addrs, err := iface.Addrs()
			if err != nil {
				return nil
			}
			for _, addr := range addrs {
				ip = append(ip, addr.String())
			}
		}
	}
	return ip
}

func GetAllIP() []string {
	ip := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ip = append(ip, addr.String())
	}
	return ip
}

func GetAllIPWithName() map[string]string {
	ips := make(map[string]string)
	list, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, iface := range list {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil
		}
		for j, addr := range addrs {
			ips[iface.Name+"-"+strconv.Itoa(j)] = addr.String()
		}
	}
	return ips
}
