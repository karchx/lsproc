package client

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var ports = "1-65535"

type ErrorCommand struct {
	message string
	command string
}

func (r *ErrorCommand) Error() string {
	return fmt.Sprintf("message %s: command %s", r.message, r.command)
}

func ListenService() {
	/*var (
		out []byte
		err error
	)

	out, err = exec.Command("/bin/bash", "-c", "cat /proc/net/tcp | awk '{ print $2 }'").Output()
	if err != nil {
		log.Fatal(err)
	}

	pidHex := strings.Split(string(out), ":")[1]
	// valor := strings.Split("8001A8C0:B8FA", ":")[1]
	fmt.Println(pidHex)*/
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pR := processRange(ctx, ports)
	sP := scanPorts(ctx, pR)
	var portsCtx []int
	for port := range sP {
		if port != 0 {
			portsCtx = append(portsCtx, port)
		}
	}
}

// RunCommand run the command with the path of the configuration file for the application
func RunCommand(command, path string) (string, error) {
	var m string = ""
	commandWithParams := strings.Split(command, " ")

	if len(commandWithParams) > 1 {
		cmd := exec.Command(commandWithParams[0], commandWithParams[1])
		cmd.Dir = path

		m = fmt.Sprintf("- Start")

		err := cmd.Start()
		if err != nil {
			return "", &ErrorCommand{message: err.Error(), command: command}
		}

		m = fmt.Sprintf("%s \nWe could do here something before we wait for the job to finish.", m)

		err = cmd.Wait()

		//fmt.Printf("- Command finished\\n")

		if err != nil {
			return "", &ErrorCommand{message: err.Error(), command: command}
		}
		return m, nil

	}

	return "", &ErrorCommand{message: "invalid command", command: command}
}

func decryptPort(hex string) int32 {
	decimal, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int32(decimal)
}
