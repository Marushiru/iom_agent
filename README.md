# iom_agent

## disk.go
 - device: windows下为盘符 linux下为文件系统(如sysfs rootfs)
 - mountpoint: 文件挂载点
 - fstype: 文件系统类型
 - opts: 
   - windows下rw是可读写, ro是只读,compress是这个盘石压缩卷
## mem.go
单位: Byte
 - total: 总内存
 - available: 可用内存
 - used: 已用内存
 - usedPercent: 已用内存百分百
 - free: 内核中获取的可用内存,没啥用, 获取可用内存用available
 
//Linux下才有的swap信息
 - swapTotal: swap分区大小
 - swapFree: swap分区可用