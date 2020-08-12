package main

import (
	"flag"
	"fmt"
	"log"
    "os"
	"os/exec"
	"runtime"
	"syscall"
)

func get_instruction_count(args []string) (count uint64) {
	// Prepare the command to be run and make it call `ptrace(PTRACE_TRACEME)`
	cmd := exec.Command(args[0])
	cmd.Args = args
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	// Wire the goroutine to its current thread before using `ptrace`
	runtime.LockOSThread()

	// Start the subprocess
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Loop and increment our counter until we exit
	var ws syscall.WaitStatus
	for {
		if _, err := syscall.Wait4(cmd.Process.Pid, &ws, syscall.WALL, nil); err != nil {
			log.Fatal(err)
		}
		if ws.Exited() {
			break
		} else if ws.Stopped() {
			if err := syscall.PtraceSingleStep(cmd.Process.Pid); err != nil {
				log.Fatal(err)
			}
			count += 1
		}
	}
	return
}

func main() {
	flag.Parse()

    if len(flag.Args()) == 0 {
        fmt.Printf("usage: %v elffile [args...]\n", os.Args[0])
        os.Exit(1)
    }

	fmt.Println(get_instruction_count(flag.Args()))
}
