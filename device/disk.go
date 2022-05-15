package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"runtime"
	"strconv"
	"strings"
)

func DiskInfo() map[string]any {
	/*
		device: windows下为盘符 linux下为文件系统(如sysfs rootfs)
		mountpoint: 文件挂载点
		fstype: 文件系统类型
		opts:
			windows下rw是可读写 ro是只读 compress是这个盘石压缩卷
	*/
	var diskInfoMap map[string]any = make(map[string]any)
	partitionsStat, err := disk.Partitions(true)
	if err == nil {
		diskInfoMap["partitionsStat"] = partitionsStat
	} else {
		diskInfoMap["partitionsStat"] = nil
	}
	//TODO 这里用 "/" 获取不到windows的整个硬盘的大小
	//系统是windows则不返回根目录UsageStat
	usageStat, err := disk.Usage("/")
	if strings.ToLower(runtime.GOOS) == "windows" {
		err = nil
	}
	if err == nil {
		diskInfoMap["usageStat"] = usageStat
	} else {
		diskInfoMap["usageStat"] = nil
	}
	diskInfoMap["diskUsage"] = DiskUsage()
	//ioCounter, _ := disk.IOCounters()
	//diskInfoMap["IOCounter"] = ioCounter
	return diskInfoMap
}

//按磁盘分区获取容量
func DiskUsage() map[string]map[string]any {
	diskUsage := map[string]map[string]any{}
	diskPartitions := []string{}
	OSType := runtime.GOOS
	//if windows 获取Windows各个盘符
	//if linux 只返回根目录"/"的用量
	if strings.ToLower(OSType) == "windows" {
		//获取windows的分区情况
		partitions, _ := disk.Partitions(true)
		//fmt.Println(partitions)
		for _, i := range partitions {
			diskPartitions = append(diskPartitions, i.Device)
		}
		//TODO 添加fstype
		for _, i := range diskPartitions {
			usage, _ := disk.Usage(i)
			diskUsage[i] = map[string]any{}
			diskUsage[i]["total"] = usage.Total
			diskUsage[i]["used"] = usage.Used
			diskUsage[i]["free"] = usage.Free
			diskUsage[i]["usedPercent"] = uint64(usage.UsedPercent)
		}
	} else if strings.ToLower(OSType) == "linux" {
		//获取 "/" 目录下所有文件(夹)名
		diskUsage["/"] = make(map[string]any)
		usage, _ := disk.Usage("/")
		diskUsage["/"]["total"] = usage.Total
		diskUsage["/"]["used"] = usage.Used
		diskUsage["/"]["free"] = usage.Free
		diskUsage["/"]["usedPercent"] = uint64(usage.UsedPercent)

	}
	//fmt.Println(diskUsage)
	return diskUsage
}

func PrintDiskInfo() {
	diskInfo := DiskInfo()["diskUsage"]
	diskInfoMap := diskInfo.(map[string]map[string]any)
	//fmt.Println(diskInfoMap)
	for k, v := range diskInfoMap {
		total := v["total"].(uint64)
		free := v["free"].(uint64)
		fmt.Println("盘符:", k, "\t总容量:", strconv.FormatFloat(float64(total)/1024/1024/1024, 'f', 2, 32), "GB", "\t剩余容量:", strconv.FormatFloat(float64(free)/1024/1024/1024, 'f', 2, 32), "GB")
	}
	partition := DiskInfo()["partitionsStat"]
	partitionSlice := partition.([]disk.PartitionStat)
	for _, de := range partitionSlice {
		fmt.Println("盘符:", de.Device, "\t文件系统: ", de.Fstype)
	}
}
