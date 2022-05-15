package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

func NetInfo() {
	//类似windows下的netstats命令返回的格式
	//fmt.Println(net.Connections("all"))

	//fmt.Println(net.ConnectionsPidWithoutUids("all", 5836))

	//获得所有网络接口的名字和状态之类的
	fmt.Println(net.Interfaces())

	//获得所有网络接口的名字和状态之类的
	//fmt.Println(net.ConntrackStats(fals))
}

//返回所有网络接口信息
func NetInterfacesInfo() net.InterfaceStatList {
	list, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	return list
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
