package device

import (
	"github.com/shirou/gopsutil/v3/disk"
	"path/filepath"
	"runtime"
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
	usageStat, err := disk.Usage("/")
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
func DiskUsage() map[string]map[string]int {
	diskUsage := make(map[string]map[string]int)
	diskPartitions := []string{}
	OSType := runtime.GOOS
	//if windows 获取Windows各个盘符
	//if linux 获取linux根目录下的文件夹
	if strings.ToLower(OSType) == "windows" {
		//获取windows的分区情况
		partitions, _ := disk.Partitions(true)
		//fmt.Println(partitions)
		for _, i := range partitions {
			diskPartitions = append(diskPartitions, i.Device)
		}
		for _, i := range diskPartitions {
			usage, _ := disk.Usage(i)
			diskUsage[i] = make(map[string]int)
			diskUsage[i]["total"] = int(usage.Total)
			diskUsage[i]["used"] = int(usage.Used)
			diskUsage[i]["free"] = int(usage.Free)
			diskUsage[i]["usedPercent"] = int(usage.UsedPercent)
		}
	} else if strings.ToLower(OSType) == "linux" {
		//获取 "/" 目录下所有文件(夹)名
		diskPartitions, _ = filepath.Glob("/*")
		for _, i := range diskPartitions {
			usage, _ := disk.Usage(i)
			diskUsage[i] = make(map[string]int)
			diskUsage[i]["total"] = int(usage.Total)
			diskUsage[i]["used"] = int(usage.Used)
			diskUsage[i]["free"] = int(usage.Free)
			diskUsage[i]["usedPercent"] = int(usage.UsedPercent)
		}
	}
	//fmt.Println(diskUsage)
	return diskUsage
}
