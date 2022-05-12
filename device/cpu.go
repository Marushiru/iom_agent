package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

func CpuInfo() map[string]any {

	var cpuInfoMap map[string]any = make(map[string]any)

	//获取Cpu实体核心数量
	physicalCores, err := cpu.Counts(false)
	if err != nil {
		fmt.Println(err)
	}
	cpuInfoMap["physicalCores"] = physicalCores

	//获取Cpu虚拟核心数量
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		fmt.Println(err)
	}
	cpuInfoMap["logicalCores"] = logicalCores

	//获取Cpus的信息
	infoStats, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	cpuInfoMap["infoStats"] = infoStats

	//获取Cpu时间???
	timeStats, err := cpu.Times(true)
	if err != nil {
		fmt.Println(err)
	}

	cpuInfoMap["timeStats"] = timeStats
	fmt.Println(cpuInfoMap)
	return cpuInfoMap
}

//TODO
func cpuUsagePercentage(duration time.Duration) {
	//获取Cpu使用百分比
	//其中interval单位是1纳秒
	for {
		perPrecent, _ := cpu.Percent(duration, true)
		fmt.Println(perPrecent)
		time.Sleep(1000000000)
	}
}
