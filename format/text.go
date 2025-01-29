package text

import (
	"fmt"
	"net"
	"strings"
)

type OutputFormatter interface {
	Format(interfaces []net.Interface) string
}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

type TextFormatter struct{}

func (f *TextFormatter) Format(interfaces []net.Interface) string {
	var output []string

	for _, iface := range interfaces {
		name := iface.Name
		hw := iface.HardwareAddr
		macStr := "N/A"

		if len(hw) == 6 {
			macStr = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
				hw[0], hw[1], hw[2], hw[3], hw[4], hw[5])
		}

		interfaceOutput := fmt.Sprintf("Interface: %s\nMAC Address: %s\n", name, macStr)
		addressOutput := "Addresses:\n"

		addrs, err := iface.Addrs()
		if err != nil {
			interfaceOutput += "\tError getting addresses for " + name
		} else {
			for _, addr := range addrs {
				addressOutput += fmt.Sprintf("\t%s\n", addr.String())
			}
		}

		output = append(output, interfaceOutput+addressOutput)
	}

	return strings.Join(output, "")
}
