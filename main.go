package main

import (
	"fmt"
	"github.com/Marushiru/iom_agent/device"
)

func main() {

	//fmt.Println("==============================================MEM===============================================")
	//device.PrintMemoryInfo()
	//fmt.Println("==============================================DISK==============================================")
	//device.PrintDiskInfo()
	//fmt.Println("==============================================DOCKER============================================")
	//device.DockerInfo()
	//fmt.Println("==============================================CPU===============================================")
	//device.PrintCpuInfo()
	//fmt.Println("==============================================HOST==============================================")
	//device.PrintHostInfo()
	//fmt.Println("==============================================NET===============================================")
	//device.NetConnection(device.KIND_TCP)
	fmt.Println("==============================================PROCESS==============================================")
	//device.NetworkIORate()
	////if device.OSInfo() == "windows" {
	////	fmt.Println("==============================================WINSERVICES==============================================")
	////	device.WinServicesInfo()
	////}
	//fmt.Println("回车建退出。。。")
	//var a string
	//fmt.Scan(&a)
	fmt.Println(device.DiskInfo())
}
