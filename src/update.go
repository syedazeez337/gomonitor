package src

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func UpdateScreen(screen tcell.Screen) {
	screen.Clear()

	width, height := screen.Size()

	// get memory status
	v, _ := mem.VirtualMemory()
	memUsed := float64(v.Used) / float64(v.Total) * 100

	// Draw a memory bar
	drawBar(screen, 1, 1, width-2, "Memory", memUsed,
		fmt.Sprintf("Total: %.2fGB Used: %.2fGB Free: %.2fGB (%.1f%%)",
			float64(v.Total)/1024/1024/1024,
			float64(v.Used)/1024/1024/1024,
			float64(v.Free)/1024/1024/1024,
			memUsed))

	// Get and display CPU usage
	c, _ := cpu.Percent(0, false)
	cpuUsed := c[0]
	drawBar(screen, 1, 3, width-2, "CPU", cpuUsed,
		fmt.Sprintf("Usage: %.1f%%", cpuUsed))

	// Get and display network stats
	n, _ := net.IOCounters(false)
	netStats := fmt.Sprintf("Network - Received: %.2fMB (%d pkts) Sent: %.2fMB (%d pkts)",
		float64(n[0].BytesRecv)/1024/1024,
		n[0].PacketsRecv,
		float64(n[0].BytesSent)/1024/1024,
		n[0].PacketsSent)
	drawText(screen, 1, 5, netStats, tcell.StyleDefault.Foreground(tcell.ColorGreen))

	// Get and display process information
	processes := GetProcessInfo()
	drawProcessTable(screen, 1, 7, width/2-1, height-7, "Top Memory Usage", SortByMemory(processes))
	drawProcessTable(screen, width/2+1, 7, width/2-2, height-7, "Top CPU Usage", SortByCPU(processes))

	screen.Show()
}

// text display
func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

// draw memory bar function
func drawBar(screen tcell.Screen, x, y, width int, label string, value float64, stats string) {
	barWidth := width - 2
	filled := int(float64(barWidth) * value / 100)

	drawText(screen, x, y, label+": "+stats, tcell.StyleDefault.Foreground(tcell.ColorYellow))

	screen.SetContent(x, y+1, '[', nil, tcell.StyleDefault)
	for i := 0; i < barWidth; i++ {
		char := ' '
		style := tcell.StyleDefault.Background(tcell.ColorDarkGray)
		if i < filled {
			style = tcell.StyleDefault.Background(tcell.ColorGreen)
		}
		screen.SetContent(x+1+i, y+1, char, nil, style)
	}
	screen.SetContent(x+barWidth+1, y+1, ']', nil, tcell.StyleDefault)
}

func drawProcessTable(screen tcell.Screen, x, y, width, height int, title string, processes []ProcessInfo) {
	// Draw table title in yellow
	drawText(screen, x, y, title, tcell.StyleDefault.Foreground(tcell.ColorYellow))

	// Draw column headers in green
	drawText(screen, x, y+1, fmt.Sprintf("%-6s %-20s %-10s %-10s",
		"PID", "Name", "CPU%", "Memory"),
		tcell.StyleDefault.Foreground(tcell.ColorGreen))

	// Draw each process row
	for i, p := range processes {
		if i >= height-3 { // Leave space for title and headers
			break
		}
		drawText(screen, x, y+2+i, fmt.Sprintf("%-6d %-20s %-10.1f %-10.1f",
			p.PID, truncateString(p.Name, 20), p.CPU, p.Memory),
			tcell.StyleDefault)
	}
}

// Helper to prevent long process names from breaking the layout
func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}
