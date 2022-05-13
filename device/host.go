package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

func HostInfo() {
	//主机信息
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hostInfo)
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
	fmt.Println(sensorsTemperatures)
}
