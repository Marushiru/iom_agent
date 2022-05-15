package device

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/docker"
)

func DockerInfo() {
	cGroupDockerStat, _ := docker.GetDockerIDList()
	fmt.Println(cGroupDockerStat)
}
