package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

func main() {
	//a := device.MemoryInfo()
	//c, _ := disk.Usage("C:")
	//d, _ := disk.Usage("D:")
	e, _ := disk.Usage("E:")
	all, _ := disk.Usage("\\")
	//fmt.Println(c.Total)
	//fmt.Println(d.Total)
	fmt.Println(e.Total)
	//total := c.Total + d.Total + e.Total
	//fmt.Println(total)
	fmt.Println(all.Total)
	//fmt.Println(runtime.GOOS)
	//b := device.DiskInfo()
	//fmt.Println(b)
	//device.DiskUsage()
}
