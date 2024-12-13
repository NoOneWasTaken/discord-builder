package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func LogStdoutStream(command *exec.Cmd) {
	commandStdOut, _ := command.StdoutPipe()
	commandStdErr, _ := command.StderrPipe()

	if err := command.Start(); err != nil {
		log.Fatal(err)
	}

	commandOutScan := bufio.NewScanner(commandStdOut)
	go func() {
		for commandOutScan.Scan() {
			fmt.Println(commandOutScan.Text())
		}
	}()

	commandErrScan := bufio.NewScanner(commandStdErr)
	go func() {
		for commandErrScan.Scan() {
			fmt.Println(commandErrScan.Text())
		}
	}()

	if err := command.Wait(); err != nil {
		log.Fatal(err)
	}
}
