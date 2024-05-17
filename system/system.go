package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
    MemTotal     int
    MemFree      int
    MemAvailable int
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

func parseLine(raw string) (key string, value int) {
    text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
    keyValue := strings.Split(text, ":")
    return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
    if raw == "" {
        return 0
    }
    res, err := strconv.Atoi(raw)
    if err != nil {
        panic(err)
    }
    return res
}