package main

import (
	"fmt"
	"net"
	"strings"
)

// 检查IP地址是否属于局域网范围
func isPrivateIP(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		// 检查是否在192.168.0.0/16范围内
		if ip4[0] == 192 && ip4[1] == 168 {
			return true
		}
		// 检查是否在172.16.0.0/12范围内
		if ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 {
			return true
		}
		// 检查是否在10.0.0.0/8范围内
		if ip4[0] == 10 {
			return true
		}
	}
	return false
}

func getIp() string {

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {

		// 筛选出可用于局域网通信的有效物理接口
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagBroadcast == 0 ||
			iface.Flags&net.FlagLoopback != 0 {
			fmt.Println("Skipping interface 1:", iface.Name)
			continue
		}

		// 过滤虚拟/容器接口
		if strings.Contains(iface.Name, "br-") ||
			strings.Contains(iface.Name, "veth") ||
			strings.Contains(iface.Name, "vEthernet") ||
			strings.Contains(iface.Name, "VMware") {
			fmt.Println("Skipping interface 2:", iface.Name)
			continue // 跳过名称包含关键词的接口
		}

		// 获取接口的地址列表
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// 遍历地址列表
		for _, addr := range addrs {
			// 检查地址类型是否为IP地址
			ipNet, ok := addr.(*net.IPNet)
			if ok {

				if isPrivateIP(ipNet.IP) {
					ipStr := ipNet.IP.String()
					return ipStr
				}
			}
		}
	}
	return ""
}
