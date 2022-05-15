package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
)

type proc process.Process

func ProcessInfo() {

	//只返回了Pids 啥都没
	fmt.Println(process.Pids())
	//只返回了Process对象，对象只包含了pid 啥都没
	fmt.Println(process.Processes())

	pro, _ := process.NewProcess(1484)
	//返回进程名
	fmt.Println(pro.Name())
	//Cpu占用百分比
	fmt.Println(pro.CPUPercent())
	//进程创建时间，13位时间戳
	fmt.Println(pro.CreateTime())
	//返回Exe文件目录
	fmt.Println(pro.Exe())
	//进程工作路径？
	fmt.Println(pro.Cwd())
	//返回应用网络流量使用情况？
	fmt.Println(pro.IOCounters())
	//返回进程是否还在运行
	fmt.Println(pro.IsRunning())
	//进程位置和其命令
	fmt.Println(pro.Cmdline())
	//进程位置和其命令 拆分成切片
	fmt.Println(pro.CmdlineSlice())
	//内存占用比例
	fmt.Println(pro.MemoryPercent())
	//返回父进程的id
	fmt.Println(pro.Ppid())
	//子进程？
	fmt.Println(pro.Children())
	//返回父进程，调用newprocess大概？
	fmt.Println(pro.Parent())

	//进程相关的网络请求
	fmt.Println(pro.Connections())

	//相关系统环境变量
	fmt.Println(pro.Environ())
	//返回应用是否在前台。win没实现
	fmt.Println(pro.Foreground())
	//返回IO优先级？？win没有实现
	fmt.Println(pro.IOnice())

	//内存表WIN没实现
	fmt.Println(pro.MemoryMaps(false))
	fmt.Println(pro.MemoryMaps(true))

	//是否后台，win没实现
	fmt.Println(pro.Background())
	//返回进程优先级
	fmt.Println(pro.Nice())

	//返回应用groupId，win没实现
	fmt.Println(pro.Gids())
	//返回应用groupId(包括补充组？)，win没实现
	fmt.Println(pro.Groups())
	//windows没有实现，返回进程的uid
	fmt.Println(pro.Uids())

	//返回资源限制？win没有实现
	fmt.Println(pro.Rlimit())
	//返回终端》？？？win没实现
	fmt.Println(pro.Terminal())
	//返回进程状态？win没有实现
	fmt.Println(pro.Status())

	//发送继续信号
	//fmt.Println(pro.Resume())
	//发送进程暂停信号？win实现了
	//fmt.Println(pro.Suspend())
	//发送进程中止信号
	//fmt.Println(pro.Terminate())
	//杀死进程
	//fmt.Println(pro.Kill())
	//发送UNIX信号到进程  syscall.Signal自己写，win没实现大概
	//fmt.Println(pro.SendSignal())

	//返回线程组ID？
	fmt.Println(pro.Tgid())

	//返回线程
	fmt.Println(pro.Threads())
	//返回线程
	fmt.Println(pro.Times())
	//返回进程的用户名
	fmt.Println(pro.Username())

	//rss用掉的物理内存，vms用掉的虚拟内存
	fmt.Println(pro.MemoryInfo())
	//杀死RSS,VMS???win没实现
	fmt.Println(pro.MemoryInfoEx())
}
