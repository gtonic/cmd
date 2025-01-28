package formatter

import (
	"encoding/json"
	"fmt"
	"net"
)

type JSONFormatter struct{}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

func (f *JSONFormatter) Format(interfaces []net.Interface) string {
	type InterfaceInfo struct {
		Name      string   `json:"name"`
		MAC       string   `json:"mac"`
		Addresses []string `json:"addresses"`
	}

	var interfaceList []InterfaceInfo

	for _, iface := range interfaces {
		name := iface.Name
		hw := iface.HardwareAddr
		macStr := "N/A"

		if len(hw) == 6 {
			macStr = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
				hw[0], hw[1], hw[2], hw[3], hw[4], hw[5])
		}

		interfaceAddresses := make([]string, 0)

		addrs, err := iface.Addrs()
		if err != nil {
			interfaceAddresses = append(interfaceAddresses, "Error getting addresses")
		} else {
			for _, addr := range addrs {
				interfaceAddresses = append(interfaceAddresses, addr.String())
			}
		}

		interfaceList = append(interfaceList, InterfaceInfo{
			Name:      name,
			MAC:       macStr,
			Addresses: interfaceAddresses,
		})
	}

	jsonBytes, err := json.MarshalIndent(interfaceList, "", "  ")
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}
