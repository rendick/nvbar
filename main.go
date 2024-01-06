package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		// Time
		// date := time.Now().Format(time.RFC850)
		date := time.Now()

		// CPU usage
		// cpu, err := exec.Command("sh", "-c", "echo "CPU Usage: "$[100-$(vmstat 1 2|tail -1|awk '{print $15}')]"%"")

		// Memory

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

		// Connection
		connection, err := exec.Command("sh", "-c", "ping -q -c1 google.com &>/dev/null && echo online || echo offline").Output()
		if err != nil {
			panic(err)
		}

		// Keyboard layout
		keyboard, err := exec.Command("xkb-switch").Output()
		if err != nil {
			panic(err)
		}

		// BAT
		battery, err := exec.Command("cat", "/sys/class/power_supply/BAT1/capacity").Output()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s"+" | "+"%d MB / %d MB"+" | "+"%s"+"| "+"%s"+"| "+"%s\n",
			strings.Replace(string(connection), "\n", "", -1),
			(total-avail)/1024, total/1024,
			strings.Replace(string(keyboard), "\n", " ", -1),
			strings.Replace(string(battery), "\n", " ", -1),
			date.Format("2006-01-02 15:04:05"))

		time.Sleep(time.Second)

	}

}
