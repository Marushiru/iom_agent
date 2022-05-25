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

//type VirtualMemoryStat struct {
//////单位: Byte
//	// 总内存
//	Total uint64 `json:"total"`
//	// 可用内存
//	Available uint64 `json:"available"`
//  //已用内存
//	Used uint64 `json:"used"`
//  //已用内存百分比
//	UsedPercent float64 `json:"usedPercent"`
//	// 内核中获取的可用内存,没啥用, 获取可用内存用available
//	Free uint64 `json:"free"`

//  用不着的东西
//	// OS X / BSD specific numbers:
//	// http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
//	Active   uint64 `json:"active"`
//	Inactive uint64 `json:"inactive"`
//	Wired    uint64 `json:"wired"`
//
//	// FreeBSD specific numbers:
//	// https://reviews.freebsd.org/D8467
//	Laundry uint64 `json:"laundry"`
//

//  Linux下的东西 具体干啥的不清楚
//	// Linux specific numbers
//	// https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
//	// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
//	// https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
//	Buffers        uint64 `json:"buffers"`
//	Cached         uint64 `json:"cached"`
//	Writeback      uint64 `json:"writeback"`
//	Dirty          uint64 `json:"dirty"`
//	WritebackTmp   uint64 `json:"writebacktmp"`
//	Shared         uint64 `json:"shared"`
//	Slab           uint64 `json:"slab"`
//	SReclaimable   uint64 `json:"sreclaimable"`
//	SUnreclaim     uint64 `json:"sunreclaim"`
//	PageTables     uint64 `json:"pagetables"`
//	SwapCached     uint64 `json:"swapcached"`
//	CommitLimit    uint64 `json:"commitlimit"`
//	CommittedAS    uint64 `json:"committedas"`
//	HighTotal      uint64 `json:"hightotal"`
//	HighFree       uint64 `json:"highfree"`
//	LowTotal       uint64 `json:"lowtotal"`
//	LowFree        uint64 `json:"lowfree"`
//	SwapTotal      uint64 `json:"swaptotal"`
//	SwapFree       uint64 `json:"swapfree"`
//	Mapped         uint64 `json:"mapped"`
//	VMallocTotal   uint64 `json:"vmalloctotal"`
//	VMallocUsed    uint64 `json:"vmallocused"`
//	VMallocChunk   uint64 `json:"vmallocchunk"`
//	HugePagesTotal uint64 `json:"hugepagestotal"`
//	HugePagesFree  uint64 `json:"hugepagesfree"`
//	HugePageSize   uint64 `json:"hugepagesize"`
//}
