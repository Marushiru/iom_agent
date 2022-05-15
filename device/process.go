package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"strconv"
	"syscall"
)

type Proc process.Process

func ProcessInfo() {

	//只返回了Pids 啥都没
	//process.Pids()
	proc, err := process.Processes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(proc)
	fmt.Println("输入你想要查看的pid:")
	var pidSt string
	fmt.Scan(&pidSt)
	pid, _ := strconv.Atoi(pidSt)
	pro, _ := process.NewProcess(int32(pid))

	//返回进程名
	name, _ := pro.Name()
	fmt.Println("进程名：" + name)
	//返回进程是否还在运行
	isRunning, _ := pro.IsRunning()
	fmt.Println("运行状态：" + strconv.FormatBool(isRunning))

	//返回进程优先级
	nice, _ := pro.Nice()
	fmt.Println("优先级：", nice)
	//返回进程的用户名
	username, _ := pro.Username()
	fmt.Println("用户名：" + username)
	//进程工作路径
	cwd, _ := pro.Cwd()
	fmt.Println("工作路径：" + cwd)

	//返回Exe文件目录
	exe, _ := pro.Exe()
	fmt.Println("执行文件路径：" + exe)
	//进程位置和其命令
	cmd, _ := pro.Cmdline()
	fmt.Println("执行文件命令：" + cmd)

	//进程创建时间，13位时间戳
	create, _ := pro.Cmdline()
	fmt.Println("创建时间戳：" + create)

	//TODO 写一个按照时间显示的方法
	//Cpu占用百分比
	cpuPercent, _ := pro.CPUPercent()
	fmt.Println("Cpu占用率：" + strconv.FormatFloat(cpuPercent, 'f', 2, 64))
	//返回应用网络流量使用情况？
	ioCounters, _ := pro.IOCounters()
	fmt.Println("流量使用情况：", ioCounters)
	//进程使用的内存，rss用掉的物理内存，vms用掉的虚拟内存

	memoryInfo, _ := pro.MemoryInfo()
	fmt.Println("内存使用情况：", memoryInfo)
	fmt.Println(pro.MemoryInfo())

	//内存占用比例
	memoryPercent, _ := pro.MemoryPercent()
	fmt.Println("内存占用率：", memoryPercent)
	fmt.Println(pro.MemoryPercent())

	//返回父进程的id
	//ppid, _ := pro.Ppid()

	//返回父进程，返回的是process对象
	parent, _ := pro.Parent()
	fmt.Println("父进程ID：", parent)

	//子进程，返回process切片
	childrens, _ := pro.Children()
	fmt.Println("子进程ID：", childrens)

	//进程相关的网络请求
	connections, _ := pro.Connections()
	fmt.Println("线程网络连接:", connections)

}

func WinNotImplement() {

}

//继续程序
func (p *Proc) Resume() {
	p.Resume()
}

//停止程序
func (p *Proc) Suspend() {
	p.Suspend()
}

//中止程序
func (p *Proc) Terminate() {
	p.Terminate()
}

//杀掉程序
func (p *Proc) Kill() {
	p.Kill()
}

//发送需要的命令
func (p *Proc) SendSignal(signal syscall.Signal) {
	p.SendSignal(signal)
}
