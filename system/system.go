package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

type Memory struct {
	MemTotal     uint64
	MemFree      uint64
	MemAvailable uint64
}

type Disk struct {
	DiskTotal uint64
	DiskFree  uint64
	DiskAvailable uint64
	DiskType  string
}

func ReadMemoryStats() Memory {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufio.NewScanner(file)
	scanner := bufio.NewScanner(file)
	res := Memory{}
	for scanner.Scan() {
		key, value := parseLine(scanner.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value
		case "MemFree":
			res.MemFree = value
		case "MemAvailable":
			res.MemAvailable = value
		}
	}
	return res
}

func ReadDiskInfo(path string) Disk {
	var info unix.Statfs_t

	unix.Statfs(path, &info)

	res := Disk{}
	res.DiskAvailable = info.Bavail * uint64(info.Bsize)
	res.DiskTotal = info.Blocks * uint64(info.Bsize)
	res.DiskFree = info.Bfree * uint64(info.Bsize)

	return res
}

func parseLine(raw string) (key string, value uint64) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
    if keyValue[1] == "" {
        return keyValue[0], 0
    }
    strconv.ParseUint(keyValue[1], 10, 64)

	return keyValue[0], toUint(keyValue[1])
}

func toUint(raw string) uint64 {
	if raw == "" {
		return 0
	}
	res, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}
