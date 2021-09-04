package main

import (
	"encoding/json"
	"fmt"

	"github.com/philcantcode/SecurityTool/agent"
	"github.com/philcantcode/SecurityTool/device"
)

var SysConfigs = make(map[string]string)

func main() {
	SysConfigs["Version"] = "1.0"
	SysConfigs["Patch"] = "1"
	SysConfigs["Mode"] = "Interactive"

	fmt.Printf("Welcome - Version %s, Patch %s\n",
		SysConfigs["Version"],
		SysConfigs["Patch"])

	switch SysConfigs["Mode"] {
	case "Interactive":
		interactiveCLI()
	}

	fmt.Println("Goodbye")
	agent.Learning()
}

func interactiveCLI() {
	var input string

	for input != "q" {
		fmt.Print("[>] ")
		fmt.Scanf("%s\n", &input)
		RunCMD(input)
	}
}

func RunCMD(cmd string) {
	switch cmd {
	case "ip":
		PrettyPrint(device.IpInfo())
	case "os":
		PrettyPrint(device.OSInfo())
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
