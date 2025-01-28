package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, iface := range interfaces {
		name := iface.Name
		hw := iface.HardwareAddr
		var macStr string
		if len(hw) == 6 {
			macStr = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
				hw[0], hw[1], hw[2], hw[3], hw[4], hw[5])
		} else {
			macStr = "N/A"
		}

		fmt.Printf("Interface: %s\nMAC Address: %s\n", name, macStr)

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Error getting addresses for %s: %v\n", name, err)
			continue
		}
		for _, addr := range addrs {
			address := addr.String()
			fmt.Printf("\t%s\n", address)
		}
	}
}
