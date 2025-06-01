package helper

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

type Data struct {
	Page int   `json:"page"`
	Data []any `json:"data"`
}

func RequestData() (*Data, error) {

	p, _ := process.Pids()
	h, _ := host.Info()
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(time.Second, false)

	mbMemory := v.Total / 1024 / 1024
	freeMemory := v.Free / 1024 / 1024

	cpuModel, _ := cpu.Info()

	numProcess := len(p)
	platInformation := h.Platform
	hostName := h.Hostname

	res := &Data{
		Page: 1,
		Data: []any{runtime.GOOS, platInformation, hostName, numProcess, "test vm", mbMemory, freeMemory, fmt.Sprintf("%.2f", c), cpuModel[0].ModelName, cpuModel[0].Model, cpuModel[0].Mhz},
	}

	return res, nil

}
