package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

//type InfoStat struct {
//	Hostname             string `json:"hostname"`
//	Uptime               uint64 `json:"uptime"`
//	BootTime             uint64 `json:"bootTime"`
//	Procs                uint64 `json:"procs"`           // number of processes
//	OS                   string `json:"os"`              // ex: freebsd, linux
//	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
//	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
//	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
//	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
//	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
//	VirtualizationSystem string `json:"virtualizationSystem"`
//	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
//	HostID               string `json:"hostid"`             // ex: uuid
//}
func HostInfo() {
	//主机信息
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("主机信息:", hostInfo)
	TemperatureInfo()
}

func OSInfo() string {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	return hostInfo.OS
}

func TemperatureInfo() {
	//传感器温度切片
	sensorsTemperatures, err := host.SensorsTemperatures()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("传感器温度信息:", sensorsTemperatures)
}
