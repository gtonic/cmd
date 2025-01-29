package main

import (
	"flag"
	"fmt"
	"net"

	json "github.com/gtonic/tool/format/json"
	text "github.com/gtonic/tool/format/text"
)

var (
	jsonOutput = flag.Bool("json", false, "output in JSON format")
)

func main() {
	flag.Parse()

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var formatter OutputFormatter

	if *jsonOutput {
		formatter = json.NewJSONFormatter()
	} else {
		formatter = text.NewTextFormatter()
	}

	output := formatter.Format(interfaces)

	fmt.Println(output)
}
