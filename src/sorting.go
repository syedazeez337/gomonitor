package src

import "sort"

func SortByMemory(processes []ProcessInfo) []ProcessInfo {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].MemSize > processes[j].MemSize
	})

	if len(processes) > 10 {
		processes = processes[:10]
	}
	return processes
}

func SortByCPU(processes []ProcessInfo) []ProcessInfo {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPU > processes[j].CPU
	})

	if len(processes) > 10 {
		processes = processes[:10]
	}
	return processes
}