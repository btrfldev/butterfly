package main

import (
	"fmt"

	"github.com/iamsoloma/butterfly/system"
)

func main() {
	mem := system.ReadMemoryStats()
	disk := system.ReadDiskInfo("./")

	fmt.Println("In MegaBytes:")
	fmt.Printf("MemTotal: %d | MemFree: %d | MemAvailable: %d\n", mem.MemTotal/1024, mem.MemFree/1024, mem.MemAvailable/1024)
	fmt.Println("In GigaBytes:")
	fmt.Printf("DiskTotal: %d | DiskFree: %d | DiskAvailable: %d\n", disk.DiskTotal/1024/1024, disk.DiskFree/1024/1024, disk.DiskAvailable/1024/1024)
}
