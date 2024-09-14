package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

type Memory struct {
	TotalMem     uint64
	FreeMem      uint64
	AvailableMem uint64
}

type Disk struct {
	TotalDisk     uint64
	FreeDisk      uint64
	AvailableDisk uint64
	DiskType      string
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
			res.TotalMem = value
		case "MemFree":
			res.FreeMem = value
		case "MemAvailable":
			res.AvailableMem = value
		}
	}
	return res
}

func ReadDiskInfo(path string) Disk {
	var info unix.Statfs_t

	unix.Statfs(path, &info)

	res := Disk{}
	res.AvailableDisk = info.Bavail * uint64(info.Bsize)
	res.TotalDisk = info.Blocks * uint64(info.Bsize)
	res.FreeDisk = info.Bfree * uint64(info.Bsize)

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
