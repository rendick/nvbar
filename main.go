package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var DateOutput string
var MemoryOutput string
var NetworkOutput string
var KeyboardOutput string
var BatteryOutput string

func Date() {
	DateOutput = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
}

func Memory() {
	avail_memory, err_avail := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemAvailable:' | awk '{print $2}'").Output()
	if err_avail != nil {
		os.Exit(0)
	}
	total_memory, err_total := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemTotal:' | awk '{print $2}'").Output()
	if err_total != nil {
		os.Exit(0)
	}

	avail_convert := strings.TrimSpace(string(avail_memory))
	avail, err_convert := strconv.Atoi(avail_convert)
	if err_convert != nil {
		panic(err_avail)
	}

	total_convert := strings.TrimSpace(string(total_memory))
	total, err_convert_total := strconv.Atoi(total_convert)
	if err_convert_total != nil {
		panic(err_total)
	}

	MemoryOutput = fmt.Sprintf("%d MB / %d MB", (total-avail)/1024, total/1024)
}

func Network() {
	network, err := exec.Command("sh", "-c", "ping -q -c1 google.com &>/dev/null && echo online || echo offline").Output()
	if err != nil {
		panic(err)
	}
	NetworkOutput = fmt.Sprintf("%s", strings.Replace(string(network), "\n", "", -1))
}

func Keyboard() {
	keyboard, err := exec.Command("xkb-switch").Output()
	if err != nil {
		panic(err)
	}
	KeyboardOutput = fmt.Sprintf("%s", strings.Replace(string(keyboard), "\n", "", -1))
}

func Battery() {
	batteryDir, err := os.Open("/sys/class/power_supply/")
	if err != nil {
		panic(err)
	}
	batteryFiles, err := batteryDir.Readdir(-1)
	batteryDir.Close()
	if err != nil {
		panic(err)
	}
	for _, file := range batteryFiles {
		if strings.HasPrefix(file.Name(), "BAT") {

			batteryCmd := exec.Command("cat", fmt.Sprintf("/sys/class/power_supply/%s/capacity", file.Name()))
			batteryOutputBytes, err := batteryCmd.Output()
			if err != nil {
				panic(err)
			}
			BatteryOutput = strings.TrimSpace(string(batteryOutputBytes)) + " "
		}
	}
}

func main() {
	for {
		Network()
		Memory()
		Keyboard()
		Battery()
		Date()
		fmt.Printf("%s"+" | "+"%s"+" | "+"%s"+" | "+"%s"+" | "+"%s\n",
			NetworkOutput,
			MemoryOutput,
			KeyboardOutput,
			strings.TrimSpace(BatteryOutput),
			DateOutput)

		time.Sleep(time.Second)
	}
}
