package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"strconv"
	"time"
)

func PrintCpuInfo() {

	PhysicalCores()
	fmt.Println()
	LogicalCores()
	fmt.Println()
	CpuInfo()
	fmt.Println()
	CpuUsagePercentage(1000000000)
	fmt.Println()
}

type CpuBaseInfo struct {
	PhysicalCores int
	LogicalCores  int
	CpuInfoStats  []cpu.InfoStat
}

func NewCpuBaseInfo() *CpuBaseInfo {
	return &CpuBaseInfo{
		PhysicalCores: PhysicalCores(),
		LogicalCores:  LogicalCores(),
		CpuInfoStats:  CpuInfo(),
	}
}

func PhysicalCores() int {
	//获取Cpu实体核心数量
	physicalCores, err := cpu.Counts(false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cpu实体核心数量: ", physicalCores, "核")
	return physicalCores
}

func LogicalCores() int {
	//获取Cpu虚拟核心数量
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cpu虚拟核心数量: ", logicalCores, "核")
	return logicalCores
}

//type InfoStat struct {
//	CPU        int32    `json:"cpu"` 		 Cpu基础信息
//	VendorID   string   `json:"vendorId"`	 厂家标识
//	Family     string   `json:"family"` 	 Cpu系列
//	Model      string   `json:"model"` 		 Cpu代号
//	Stepping   int32    `json:"stepping"` 	 更新版本
//	PhysicalID string   `json:"physicalId"`  物理核编号
//	CoreID     string   `json:"coreId"`		 物理Cpu标号
//	Cores      int32    `json:"cores"` 		 核心数
//	ModelName  string   `json:"modelName"`   Cpu全称
//	Mhz        float64  `json:"mhz"` 		 主频
//	CacheSize  int32    `json:"cacheSize"`   二级缓存
//	Flags      []string `json:"flags"`		 支持
//	Microcode  string   `json:"microcode"`	 微指令
//}

func CpuInfo() []cpu.InfoStat {
	//获取Cpus的信息
	infoStats, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}

	for _, j := range infoStats {
		fmt.Println("Cpu基础信息")
		fmt.Println("核心 ", j.CPU, "\t", "Cpu系列：", j.Family, "\t", "Cpu代号：", j.Model, "\t", "厂家标识：", j.VendorID, "\t")
		fmt.Println("更新版本：", j.Stepping, "\t", "物理核编号: ", j.CoreID, "\t", "核心数: ", j.Cores, "\t", "物理Cpu标号: ", j.PhysicalID, "\t")
		fmt.Println("主频:", j.Mhz, "\t", "二级缓存:", j.CacheSize, "\t", "微指令:", j.Microcode, "\t", "支持:", j.Flags, "\t")
		println("Cpu全称:", j.ModelName, "\t")
	}

	return infoStats
}

// TODO
// Cpu的使用率需要按照时间定时传给服务器
func CpuUsagePercentage(duration time.Duration) []float64 {
	//获取Cpu使用百分比
	//其中interval单位是1纳秒
	fmt.Println("核心使用率")
	perPrecent, _ := cpu.Percent(duration, true)
	for i, j := range perPrecent {
		if i%4 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Print("核心 ", i, "\t", strconv.FormatFloat(j, 'f', 2, 64), "\t\t")
	}
	return perPrecent
}
