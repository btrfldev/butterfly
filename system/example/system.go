package main

import (
	"fmt"

	"github.com/btrfldev/butterfly/system"
)

func main() {
	mem := system.ReadMemoryStats()
	disk := system.ReadDiskInfo("./")

	fmt.Println("In MegaBytes:")
	fmt.Printf("MemTotal: %d | MemFree: %d | MemAvailable: %d\n", mem.TotalMem/1024, mem.FreeMem/1024, mem.AvailableMem/1024)
	fmt.Println("In GigaBytes:")
	fmt.Printf("DiskTotal: %d | DiskFree: %d | DiskAvailable: %d\n", disk.TotalDisk/1024/1024, disk.FreeDisk/1024/1024, disk.AvailableDisk/1024/1024)
}
