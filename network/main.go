package network

import (
	"net"
	"strconv"
)

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
