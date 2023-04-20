package testmemory

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type MemInfo struct {
	Total     int
	Free      int
	Available int
}

func StringToInt(raw string) int {
	if raw == "" {
		return 0
	} else {
		res, err := strconv.Atoi(raw)
		if err != nil {
			panic(err)
		}
		return res
	}
}

func ParseRaw(raw string) (key string, value int) {
	str_replace := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	str_split := strings.Split(str_replace, ":")
	return str_split[0], StringToInt(str_split[1])
}

func ReadMemInfo() MemInfo {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var meminfo MemInfo
	for scanner.Scan() {
		key, value := ParseRaw(scanner.Text())
		switch key {
			case "MemTotal":
				meminfo.Total = value
			case "MemFree":
				meminfo.Free = value
			case "MemAvailable":
				meminfo.Available = value
		}
	}
	return meminfo
}
