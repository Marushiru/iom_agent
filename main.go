package main

import (
	"fmt"
	"github.com/Marushiru/iom_agent/device"
)

func main() {
	//a := device.MemoryInfo()
	b := device.DiskInfo()
	//c, _ := disk.Usage("\\\\wsl$\\Ubuntu-18.04")
	//fmt.Println(c)
	fmt.Println(b)
}
