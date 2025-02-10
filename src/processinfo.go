package src

import "github.com/shirou/gopsutil/v3/process"

func GetProcessInfo() []ProcessInfo {
	processes, _ := process.Processes()
	var processInfo []ProcessInfo

	for _, p := range processes {
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()
		memInfo, _ := p.MemoryInfo()

		if memInfo != nil {
			processInfo = append(processInfo, ProcessInfo{
				PID: p.Pid,
				Name: name,
				CPU: cpu,
				Memory: float64(mem),
				MemSize: memInfo.RSS,
			})
		}
	}
	return processInfo
}