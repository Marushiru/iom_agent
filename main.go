package main

import (
	"fmt"
	"github.com/Marushiru/iom_agent/device"
)

func main() {
	fmt.Println("==============================================CPU==============================================")
	device.CpuInfo()
	fmt.Println("==============================================HOST==============================================")
	device.HostInfo()
	fmt.Println("==============================================NET==============================================")
	device.NetInfo()
	fmt.Println("==============================================PROCESS==============================================")
	device.ProcessInfo()
	if device.OSInfo() == "windows" {
		fmt.Println("==============================================WINSERVICES==============================================")
		device.WinServicesInfo()
	}
}
