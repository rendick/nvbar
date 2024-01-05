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
	for i := 0; i < 20; i++ {
		// s, _ := f.Marshal(obj)
		// fmt.Println(string(s))
		// Time
		date := time.Now().Format(time.RFC850)

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

		fmt.Printf("%d MB / %d MB ", (total-avail)/1024, total/1024)

		fmt.Printf("%s\n", date)

		time.Sleep(time.Second)

	}

}
