package main

import (
	"flag"
	"fmt"
	"os"
    "github.com/rajat8rj/system-info-cli/internal/cpu"
    "github.com/rajat8rj/system-info-cli/internal/mem"
    "github.com/rajat8rj/system-info-cli/internal/disk"
)

func main() {
	cpuFlag := flag.Bool("cpu", false, "Show CPU info")
	memFlag := flag.Bool("mem", false, "Show Memory info")
	diskFlag := flag.Bool("disk", false, "Show Disk info")
	flag.Parse()

	if *cpuFlag {
		fmt.Println("CPU Info",cpu.GetCPUInfo())
	}
	if *memFlag {
		fmt.Println(mem.GetMemInfo())
	}
	if *diskFlag {
		fmt.Println(disk.GetDiskInfo())
	}
	if !*cpuFlag && !*memFlag && !*diskFlag {
		fmt.Println("Usage: ./main --cpu | --mem | --disk")
		os.Exit(1)
	}
}