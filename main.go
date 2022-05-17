package main

import (
	"fmt"
	"github.com/Marushiru/iom_agent/device"
)

func main() {
	fmt.Println("==============================================MEM===============================================")
	device.PrintMemoryInfo()
	fmt.Println("==============================================DISK==============================================")
	device.PrintDiskInfo()
	fmt.Println("==============================================DOCKER============================================")
	device.DockerInfo()
	fmt.Println("==============================================CPU===============================================")
	device.PrintCpuInfo()
	fmt.Println("==============================================HOST==============================================")
	device.PrintHostInfo()
	fmt.Println("==============================================NET===============================================")
	device.PrintNetInfo()
	//fmt.Println("==============================================PROCESS==============================================")
	//device.ProcessInfo()
	//if device.OSInfo() == "windows" {
	//	fmt.Println("==============================================WINSERVICES==============================================")
	//	device.WinServicesInfo()
	//}

}
