package presenter

import (
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

func MainTable() {

	p, _ := process.Pids()
	h, _ := host.Info()

	numProcess := len(p)
	platInformation := h.Platform
	hostName := h.Hostname

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

func HardwareTable() {

	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(time.Second, false)

	mbMemory := v.Total / 1024 / 1024
	freeMemory := v.Free / 1024 / 1024

	fmt.Printf("======== Hardware Stats ========\n")
	t2 := table.NewWriter()
	t2.SetOutputMirror(os.Stdout)
	t2.AppendHeader(table.Row{"VM Name", "Total Memory", "Free Memory", "Cpu Usage"})
	t2.AppendRows([]table.Row{
		{"test vm", mbMemory, freeMemory, fmt.Sprintf("%.2f", c)},
	})
	t2.Render()
}

func CpuTable() {

	cpuModel, _ := cpu.Info()

	fmt.Printf("======== CPU Stats ========\n")

	t3 := table.NewWriter()
	t3.SetOutputMirror(os.Stdout)
	t3.AppendHeader(table.Row{"Model Name", "Model", "Speed"})
	t3.AppendRows([]table.Row{
		{cpuModel[0].ModelName, cpuModel[0].Model, cpuModel[0].Mhz},
	})
	t3.Pager()
	t3.Render()
}
