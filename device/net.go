package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

//Family
//AF_UNSPEC  = 0
//AF_UNIX    = 1     本地通信
//AF_INET    = 2
//AF_INET6   = 23
//AF_NETBIOS = 17

//Type
//SOCK_STREAM    = 1
//SOCK_DGRAM     = 2
//SOCK_RAW       = 3
//SOCK_SEQPACKET = 5

//type ConnectionStat struct {
//	Fd     uint32  `json:"fd"`         文件描述符？？？
//	Family uint32  `json:"family"`     套接口地址结构的类型？？限制能发什么或者能收什么
//	Type   uint32  `json:"type"`       socket类型
//	Laddr  Addr    `json:"localaddr"`  本地地址
//	Raddr  Addr    `json:"remoteaddr"` 远程地址
//	Status string  `json:"status"`     状态 listen之类的
//	Uids   []int32 `json:"uids"`       用户？谁打开的这个程序？
//	Pid    int32   `json:"pid"`        程序Pid
//}

const (
	//kind
	KIND_ALL  = "all"  //代表 TCP 协议，其基于的 IP 协议的版本根据参数address的值自适应。
	KIND_TCP  = "tcp"  //代表 TCP 协议，其基于的 IP 协议的版本根据参数address的值自适应。
	KIND_TCP4 = "tcp4" //代表基于 IP 协议第四版的 TCP 协议。
	KIND_TCP6 = "tcp6" //代表基于 IP 协议第六版的 TCP 协议。
	KIND_UDP  = "udp"  //代表 UDP 协议，其基于的 IP 协议的版本根据参数address的值自适应。
	KIND_UDP4 = "udp4" //代表基于 IP 协议第四版的 UDP 协议。
	KIND_UDP6 = "udp6" //代表基于 IP 协议第六版的 UDP 协议。
	KIND_UNIX = "unix" //代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_STREAM 为 socket 类型。
	//"unixgram"：代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_DGRAM 为 socket 类型。
	//"unixpacket"：代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_SEQPACKET 为 socket 类型。
)

//返回机子上所有kind协议的连接
//暂时不知道用不用得着 留着把
func NetConnection(kind string) []net.ConnectionStat {
	conns, err := net.Connections(kind)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conns)
	return conns
}

//type InterfaceStat struct {
//	Index        int             `json:"index"`        // 大概是个ID之类的东西
//	MTU          int             `json:"mtu"`          // 最大传输单元   maximum transmission unit
//	Name         string          `json:"name"`         // 接口名        e.g., "en0", "lo0", "eth0.100"
//	HardwareAddr string          `json:"hardwareaddr"` // MAC地址       IEEE MAC-48, EUI-48 and EUI-64 form
//	Flags        []string        `json:"flags"`        // 标识？大概是支持的功能     e.g., FlagUp, FlagLoopback, FlagMulticast
//	Addrs        []InterfaceAddr `json:"addrs"`        // 其他地址
//}
//这个接口返回所有固定信息

func NewNetInterfacesInfo() net.InterfaceStatList {
	//返回所有网络接口信息
	fmt.Println("网络接口信息")
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, j := range interfaces {
		fmt.Println("ID：", j.Index, "\t\t")
		fmt.Println("接口名：", j.Name, "\t")
		fmt.Println("最大传输单元：", j.MTU, "\t")
		fmt.Println("标识：", j.Flags, "\t")
		fmt.Println("MAC地址：", j.HardwareAddr, "\t\t")
		fmt.Println("其他地址：", j.Addrs, "\t\t")
		fmt.Println()
	}
	return interfaces
}

//type IOCountersStat struct {
//	Name        string `json:"name"`        //接口名   interface name
//	BytesSent   uint64 `json:"bytesSent"`   //发送字节 number of bytes sent
//	BytesRecv   uint64 `json:"bytesRecv"`   //接收字节 number of bytes received
//	PacketsSent uint64 `json:"packetsSent"` //发送包   number of packets sent
//	PacketsRecv uint64 `json:"packetsRecv"` //接收包   number of packets received
//	Errin       uint64 `json:"errin"`       //发送错误 total number of errors while receiving
//	Errout      uint64 `json:"errout"`      //接受错误 total number of errors while sending
//	Dropin      uint64 `json:"dropin"`      //发送丢包 total number of incoming packets which were dropped
//	Dropout     uint64 `json:"dropout"`     //接收丢包 total number of outgoing packets which were dropped (always 0 on OSX and BSD)
//	Fifoin      uint64 `json:"fifoin"`      //FIFO缓存发送错误 total number of FIFO buffers errors while receiving
//	Fifoout     uint64 `json:"fifoout"`     //FIFO缓存接受错误 total number of FIFO buffers errors while sending
//}
//TODO
//这个方法也是要按照时间调用的

func IOCountersInfo() []net.IOCountersStat {
	counter, err := net.IOCounters(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("网络接口流量")
	for _, j := range counter {
		fmt.Println("接口名：", j.Name, "\t")
		fmt.Println("发送字节：", j.BytesSent, "\t\t")
		fmt.Println("接收字节：", j.BytesRecv, "\t")
		fmt.Println("发送包：", j.PacketsSent, "\t\t")
		fmt.Println("接收包：", j.PacketsRecv, "\t")
		fmt.Println("发送错误：", j.Errout, "\t\t")
		fmt.Println("接受错误：", j.Errin, "\t")
		fmt.Println("发送丢包：", j.Dropout, "\t\t")
		fmt.Println("接收丢包：", j.Dropin, "\t")
		fmt.Println("FIFO缓存发送错误：", j.Fifoout, "\t")
		fmt.Println("FIFO缓存接受错误：", j.Fifoin, "\t")
		fmt.Println()
	}
	return counter
}

type NetworkIORates struct {
	BytesSent uint64
	BytesRecv uint64
}

func NewNetworkIORates(sent uint64, recv uint64) NetworkIORates {
	return NetworkIORates{
		BytesSent: sent,
		BytesRecv: recv,
	}
}

func NetworkIORate() {
	//pernic设置为false则为计算返回所有接口的
	//           true则为分别计算
	last, err := net.IOCounters(true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(last)
	var networkIOCountMap map[string]any = make(map[string]any)

	for {
		ioCounters, err := net.IOCounters(true)
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range ioCounters {
			bytesSent := ioCounters[k].BytesSent - last[k].BytesSent
			bytesRecv := ioCounters[k].BytesRecv - last[k].BytesRecv
			networkIOCountMap[v.Name] = NewNetworkIORates(bytesSent, bytesRecv)
		}
		last = ioCounters
		fmt.Println(networkIOCountMap)

		time.Sleep(1000000000)
	}
}
