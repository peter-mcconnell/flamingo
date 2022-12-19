package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func checkCommand(command string) error {
	// check if command is in PATH
	_, err := exec.LookPath(command)
	if err != nil {
		return fmt.Errorf("'%s' command not found in PATH", command)
	}
	return nil
}

func checkProcessExists(pid int) error {
	// check if process with given PID exists
	_, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process with PID %d not found", pid)
	}
	return nil
}

func main() {
	fmt.Println("flamegraph tool")
	fmt.Println("~ just a wrapper for perf and flamegraph ~")
	// check if perf and flamegraph are installed
	if err := checkCommand("perf"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := checkCommand("flamegraph.pl"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// parse the flags
	fmt.Println("parsing flags ...")
	pidFlag := flag.Int("pid", 0, "process ID to generate flamegraph for")
	sleepFlag := flag.Int("s", 10, "sample duration in seconds")
	widthFlag := flag.Int("w", 2000, "flamegraph width")
	flag.Parse()

	if *pidFlag == 0 {
		fmt.Println("Error: --pid flag is required")
		os.Exit(1)
	}

	// check if process with given PID exists
	fmt.Println("checking process ...")
	if err := checkProcessExists(*pidFlag); err != nil {
		fmt.Printf("process %d does not exist!", pidFlag)
		os.Exit(1)
	}

	// run perf and generate a flamegraph
	fmt.Println("running perf ...")
	pidStr := strconv.Itoa(*pidFlag)
	sleepStr := strconv.Itoa(*sleepFlag)
	widthStr := strconv.Itoa(*widthFlag)
	perfOut, err := exec.Command("perf", "record", "-a", "-g", "--pid", pidStr, "sleep", sleepStr).CombinedOutput()
	fmt.Printf("got: %s\n", perfOut)
	if err != nil {
		fmt.Printf("failed when running perf\n", err)
		fmt.Println(string(perfOut))
		os.Exit(1)
	}

	fmt.Println("running flamegraph...")
	cmd := fmt.Sprintf("perf script -f | stackcollapse-perf.pl | flamegraph.pl --color=java --hash --title=Flamegraph --width=%s", widthStr)
	flamegraphOut, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(flamegraphOut))
		os.Exit(1)
	}

	// write flamegraph to stdout
	fmt.Print(string(flamegraphOut))
}
