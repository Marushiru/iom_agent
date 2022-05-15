package device

import (
	"github.com/shirou/gopsutil/v3/disk"
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
	//TODO 这里用 "/" 获取不到windows的整个硬盘的大小
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
func DiskUsage() map[string]map[string]uint64 {
	diskUsage := make(map[string]map[string]uint64)
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
			diskUsage[i] = make(map[string]uint64)
			diskUsage[i]["total"] = usage.Total
			diskUsage[i]["used"] = usage.Used
			diskUsage[i]["free"] = usage.Free
			diskUsage[i]["usedPercent"] = uint64(usage.UsedPercent)
		}
	} else if strings.ToLower(OSType) == "linux" {
		//获取 "/" 目录下所有文件(夹)名
		diskUsage["/"] = make(map[string]uint64)
		usage, _ := disk.Usage("/")
		diskUsage["/"]["total"] = usage.Total
		diskUsage["/"]["used"] = usage.Used
		diskUsage["/"]["free"] = usage.Free
		diskUsage["/"]["usedPercent"] = uint64(usage.UsedPercent)

	}
	//fmt.Println(diskUsage)
	return diskUsage
}
