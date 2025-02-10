package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func main() {

	v, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v, Used: %v, Free: %v\n", v.Total, v.Used, v.Free)
}
