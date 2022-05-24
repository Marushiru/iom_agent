package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

func PrintHostInfo() {
	TemperatureInfo()
}

//type InfoStat struct {
//	Hostname             string `json:"hostname"`        主机名字
//	Uptime               uint64 `json:"uptime"`          使用时间
//	BootTime             uint64 `json:"bootTime"`        开机时间
//	Procs                uint64 `json:"procs"`           进程数量	// number of processes
//	OS                   string `json:"os"`              操作系统	// ex: freebsd, linux
//	Platform             string `json:"platform"`        平台   		// ex: ubuntu, linuxmint
//	PlatformFamily       string `json:"platformFamily"`  平台系列？	// ex: debian, rhel
//	PlatformVersion      string `json:"platformVersion"` 平台版本	// version of the complete OS
//	KernelVersion        string `json:"kernelVersion"`   操作系统核心版本// version of the OS kernel (if available)
//	KernelArch           string `json:"kernelArch"`      操作系统核心架构// native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
//	VirtualizationSystem string `json:"virtualizationSystem"`     虚拟系统？目前没见返回过
//	VirtualizationRole   string `json:"virtualizationRole"`       虚拟角色？// guest or host
//	HostID               string `json:"hostid"`          主机ID// ex: uuid
//}

func NewHostInfo() *host.InfoStat {
	//主机信息
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("主机名称:", hostInfo.Hostname)
	fmt.Println("主机Uuid:", hostInfo.HostID)
	fmt.Println("使用时间:", hostInfo.Uptime)
	fmt.Println("开机时间:", hostInfo.BootTime)
	fmt.Println("进程数:", hostInfo.Procs)
	fmt.Println("操作系统:", hostInfo.OS)
	fmt.Println("内核架构:", hostInfo.KernelArch)
	fmt.Println("内核版本:", hostInfo.KernelVersion)
	fmt.Println("系统平台:", hostInfo.Platform)
	fmt.Println("系统版本:", hostInfo.PlatformVersion)
	fmt.Println("平台所属:", hostInfo.PlatformFamily)
	fmt.Println("虚拟系统:", hostInfo.VirtualizationSystem)
	fmt.Println("虚拟角色:", hostInfo.VirtualizationRole)
	return hostInfo
}

func OSInfo() string {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	return hostInfo.OS
}

func TemperatureInfo() []host.TemperatureStat {
	//传感器温度切片
	sensorsTemperatures, err := host.SensorsTemperatures()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("传感器温度信息:")
	for _, j := range sensorsTemperatures {
		fmt.Println("设备名：", j.SensorKey, "\t\t", "温度：", j.Temperature, "\t", "最高温度：", j.High, "\t", "告警温度：", j.Critical, "\t")
	}
	return sensorsTemperatures
}
