package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/net"
    "github.com/shirou/gopsutil/v3/process"
)



func main() {

	v, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v, Used: %v, Free: %v\n", v.Total, v.Used, v.Free)
	fmt.Println("Used Percentage:", v.UsedPercent)

	percentages, _ := cpu.Percent(0, false)
	cpuUsed := percentages[0]
	fmt.Println("CPU percentage:", cpuUsed)

	netstates, _ := net.IOCounters(false)
	stats := netstates[0]
	fmt.Println("Network statistics:", stats)

	processes, _ := process.Processes()
	for _, p := range processes {
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()
		memInfo, _ := p.MemoryInfo()

		fmt.Println(name, cpu, mem, memInfo)
	}
}
