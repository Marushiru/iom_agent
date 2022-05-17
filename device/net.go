package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

func NetInfo() {
	IOCountersInfo()
	NetInterfacesInfo()

}

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

//type InterfaceStat struct {
//	Index        int             `json:"index"`        // 大概是个ID之类的东西
//	MTU          int             `json:"mtu"`          // 最大传输单元   maximum transmission unit
//	Name         string          `json:"name"`         // 接口名        e.g., "en0", "lo0", "eth0.100"
//	HardwareAddr string          `json:"hardwareaddr"` // MAC地址       IEEE MAC-48, EUI-48 and EUI-64 form
//	Flags        []string        `json:"flags"`        // 标识？大概是支持的功能     e.g., FlagUp, FlagLoopback, FlagMulticast
//	Addrs        []InterfaceAddr `json:"addrs"`        // 其他地址
//}
//返回所有网络接口信息
func NetInterfacesInfo() net.InterfaceStatList {
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
