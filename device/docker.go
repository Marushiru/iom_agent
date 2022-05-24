package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/docker"
)

func NewDockerInfo() []docker.CgroupDockerStat {
	dockerStat, err := docker.GetDockerStat()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(dockerStat)
	for index, docker := range dockerStat {
		fmt.Println("-------DOCKER NO.", index, "-------")
		fmt.Println("containerID: ", docker.ContainerID)
		fmt.Println("docker名: ", docker.Name)
		fmt.Println("使用镜像: ", docker.Image)
		fmt.Println("状态: ", docker.Status)
		fmt.Println("是否运行中: ", docker.Running)
	}
	return dockerStat
	//dockerIDList, _ := docker.GetDockerIDList()
	//for _, id := range dockerIDList {
	//	fmt.Println(id)
	//	dockerStatByID, _ := docker.GetDockerStat()
	//	println(dockerStatByID)
	//}
}
