package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"strconv"
)

func NewMemoryInfo() *mem.VirtualMemoryStat {
	/*
		单位: Byte
		total: 总内存
		available: 可用内存
		used: 已用内存
		usedPercent: 已用内存百分百
		free: 内核中获取的可用内存,没啥用, 获取可用内存用available

		Linux下才有swap信息
		swapTotal: swap分区大小
		swapFree: swap分区可用
	*/
	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return memoryStat
}

func PrintMemoryInfo() {
	memInfo := NewMemoryInfo()
	fmt.Println("总内存:  ", strconv.FormatFloat(float64(memInfo.Total)/1024/1024, 'f', 2, 32), "MB")
	fmt.Println("已用内存: ", strconv.FormatFloat(float64(memInfo.Used)/1024/1024, 'f', 2, 32), "MB")
	fmt.Println("剩余内存: ", strconv.FormatFloat(float64(memInfo.Available)/1024/1024, 'f', 2, 32), "MB")
	fmt.Println("已用百分比: ", strconv.FormatFloat(memInfo.UsedPercent, 'f', 2, 32), "%")
}
