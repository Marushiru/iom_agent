package device

import (
	"encoding/json"
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
	//diskInfoMap["diskUsage"] = DiskUsage()
	//ioCounter, _ := disk.IOCounters()
	//diskInfoMap["IOCounter"] = ioCounter
	return diskInfoMap
}

type DiskUsage struct {
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent string
	Mountpoint  string
}

//按磁盘分区获取容量
func NewDiskUsage() map[string]DiskUsage {
	diskUsage := map[string]DiskUsage{}
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
		//TODO 添加opts
		for _, i := range diskPartitions {
			usage, _ := disk.Usage(i)

			diskUsage[i] = DiskUsage{
				Total:       usage.Total,
				Used:        usage.Used,
				Free:        usage.Free,
				UsedPercent: strconv.FormatFloat(usage.UsedPercent, 'f', 2, 64),
			}
		}
	} else if strings.ToLower(OSType) == "linux" {
		partitions, _ := disk.Partitions(true)
		for _, partition := range partitions {
			usage, _ := disk.Usage(partition.Mountpoint)

			diskUsage[partition.Device] = DiskUsage{
				Total:       usage.Total,
				Used:        usage.Used,
				Free:        usage.Free,
				UsedPercent: strconv.FormatFloat(usage.UsedPercent, 'f', 2, 64),
				Mountpoint:  partition.Mountpoint,
			}

		}

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

func PrintDiskUsage() {
	diskUsage := NewDiskUsage()
	o, _ := json.Marshal(diskUsage)
	fmt.Println(string(o))
	//for k, v := range diskUsage {
	//	fmt.Print(k, "\t")
	//	for i, j := range v {
	//		if i == "usedPercent" {
	//			fmt.Print(i, ":", j, "%\t")
	//		} else if i == "mountPoint" {
	//			fmt.Print(i, ":", j, "\t")
	//		} else {
	//			fmt.Print(i, ":", j.(uint64)/1024/1024, "MB\t")
	//		}
	//	}
	//	fmt.Println()
	//}
}
