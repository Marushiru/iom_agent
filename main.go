package main

import (
	"fmt"
	"github.com/Marushiru/iom_agent/device"
)

func main() {
	//a := device.MemoryInfo()
	b := device.DiskInfo()
	//c, _ := disk.Usage("/etc")
	//fmt.Println(c)
	//fmt.Println(runtime.GOOS)
	fmt.Println(b)
	//device.DiskUsage()
}
