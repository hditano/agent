package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

type response struct {
	Page int           `json:"page"`
	Data []interface{} `json:"vms"`
}

func main() {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			v, _ := mem.VirtualMemory()
			c, _ := cpu.Percent(time.Second, false)
			p, _ := process.Pids()
			h, _ := host.Info()

			mbMemory := v.Total / 1024 / 1024
			freeMemory := v.Free / 1024 / 1024

			numProcess := len(p)
			platInformation := h.Platform
			hostName := h.Hostname
			cpuModel, _ := cpu.Info()

			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			mainTable(platInformation, hostName, numProcess)
			hardwareTable(mbMemory, freeMemory, c[0])
			cpuTable(cpuModel[0].ModelName, cpuModel[0].Model, cpuModel[0].Mhz)

			res := &response{
				Page: 1,
				Data: []interface{}{platInformation, hostName, numProcess, mbMemory, freeMemory, c[0], cpuModel[0].ModelName, cpuModel[0].Model, cpuModel[0].Mhz},
			}
			res1B, _ := json.Marshal(res)
			fmt.Println(string(res1B))
		}
	}
}

func mainTable(platInformation string, hostName string, numProcess int) {

	fmt.Printf("======== Main Stats ========\n")
	t1 := table.NewWriter()
	t1.SetOutputMirror(os.Stdout)
	t1.AppendHeader(table.Row{"O.S", "Platform", "Hostname", "Processes Running"})
	t1.AppendRows([]table.Row{
		{runtime.GOOS, platInformation, hostName, numProcess},
	})
	t1.Pager()
	t1.Render()
}

func hardwareTable(memory uint64, freeMemory uint64, c float64) {

	fmt.Printf("======== Hardware Stats ========\n")
	t2 := table.NewWriter()
	t2.SetOutputMirror(os.Stdout)
	t2.AppendHeader(table.Row{"VM Name", "Total Memory", "Free Memory", "Cpu Usage"})
	t2.AppendRows([]table.Row{
		{"test vm", memory, freeMemory, fmt.Sprintf("%.2f", c)},
	})
	t2.Render()
}

func cpuTable(modelName string, model string, speed float64) {

	fmt.Printf("======== CPU Stats ========\n")

	t3 := table.NewWriter()
	t3.SetOutputMirror(os.Stdout)
	t3.AppendHeader(table.Row{"Model Name", "Model", "Speed"})
	t3.AppendRows([]table.Row{
		{modelName, model, speed},
	})
	t3.Pager()
	t3.Render()
}
