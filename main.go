package main

import (
	"fmt"
	"github.com/Marushiru/iom_agent/device"
)

func main() {
	a := device.MemoryInfo()
	fmt.Println(a)
}
