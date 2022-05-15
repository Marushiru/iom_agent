package device

import "github.com/shirou/gopsutil/v3/disk"

func DiskInfo() map[string]any {
	/*
		device: windows下为盘符 linux下为文件系统(如sysfs rootfs)
		mountpoint: 文件挂载点
		fstype: 文件系统类型
		opts:
			windows下rw是可读写 ro是只读 compress是这个盘石压缩卷
	*/
	var diskInfoMap map[string]any = make(map[string]any)
	partitionsStat, _ := disk.Partitions(true)
	diskInfoMap["partitionsStat"] = partitionsStat
	usageStat, _ := disk.Usage("/")
	diskInfoMap["usageStat"] = usageStat
	//ioCounter, _ := disk.IOCounters()
	//diskInfoMap["IOCounter"] = ioCounter

	return diskInfoMap
}
