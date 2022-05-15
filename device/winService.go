package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/winservices"
)

func WinServicesInfo() {

	//列出所有的service
	services, err := winservices.ListServices()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(services)
	fmt.Println("输入你想要查看的服务:")
	var name string
	fmt.Scan(&name)
	serv, err := winservices.NewService(name)
	if err != nil {
		fmt.Println(err)
	}

	//type Config struct {
	//	ServiceType      uint32
	//	StartType        uint32
	//	ErrorControl     uint32
	//	BinaryPathName   string // fully qualified path to the service binary file, can also include arguments for an auto-start service
	//	LoadOrderGroup   string
	//	TagId            uint32
	//	Dependencies     []string
	//	ServiceStartName string // name of the account under which the service should run
	//	DisplayName      string
	//	Password         string
	//	Description      string
	//	SidType          uint32 // one of SERVICE_SID_TYPE, the type of sid to use for the service
	//	DelayedAutoStart bool   // the service is started after other auto-start services are started plus a short delay
	//}
	config, err := serv.QueryServiceConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("服务配置信息：", config)

	//type ServiceStatus struct {
	//	State         svc.State       State describes service execution state (Stopped, Running and so on).
	//	Accepts       svc.Accepted    Accepted is used to describe commands accepted by the service. Note that Interrogate is always accepted.
	//	Pid           uint32
	//	Win32ExitCode uint32
	//}
	status, err := serv.QueryStatus()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("服务状态信息：", status)

}
