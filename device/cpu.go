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
	fmt.Println("Cpu实体核心数量", physicalCores)
	cpuInfoMap["physicalCores"] = physicalCores

	//获取Cpu虚拟核心数量
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cpu虚拟核心数量", logicalCores)
	cpuInfoMap["logicalCores"] = logicalCores

	//type InfoStat struct {
	//	CPU        int32    `json:"cpu"`
	//	VendorID   string   `json:"vendorId"`
	//	Family     string   `json:"family"`
	//	Model      string   `json:"model"`
	//	Stepping   int32    `json:"stepping"`
	//	PhysicalID string   `json:"physicalId"`
	//	CoreID     string   `json:"coreId"`
	//	Cores      int32    `json:"cores"`
	//	ModelName  string   `json:"modelName"`
	//	Mhz        float64  `json:"mhz"`
	//	CacheSize  int32    `json:"cacheSize"`
	//	Flags      []string `json:"flags"`
	//	Microcode  string   `json:"microcode"`
	//}
	//获取Cpus的信息
	infoStats, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cpu基础信息", infoStats)
	cpuInfoMap["infoStats"] = infoStats

	//获取Cpu时间???

	//type TimesStat struct {
	//	CPU       string  `json:"cpu"`
	//	User      float64 `json:"user"`
	//	System    float64 `json:"system"`
	//	Idle      float64 `json:"idle"`
	//	Nice      float64 `json:"nice"`
	//	Iowait    float64 `json:"iowait"`
	//	Irq       float64 `json:"irq"`
	//	Softirq   float64 `json:"softirq"`
	//	Steal     float64 `json:"steal"`
	//	Guest     float64 `json:"guest"`
	//	GuestNice float64 `json:"guestNice"`
	//}
	timeStats, err := cpu.Times(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cpu时间信息", timeStats)
	cpuInfoMap["timeStats"] = timeStats

	CpuUsagePercentage(1000000000)

	return cpuInfoMap
}

//TODO
func CpuUsagePercentage(duration time.Duration) {
	//获取Cpu使用百分比
	//其中interval单位是1纳秒
	//for {
	perPrecent, _ := cpu.Percent(duration, true)
	fmt.Println("Cpu使用率信息", perPrecent)
	//time.Sleep(1000000000)
	//}
}
