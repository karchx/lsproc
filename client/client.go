package client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var ports = "1-65535"

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

	for port := range sP {
		if strings.HasSuffix(port, ": Abierto") {
			ports := strings.Split(port, ": Abierto")
			fmt.Print(ports)
		}
	}
}

func decryptPort(hex string) int32 {
	decimal, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int32(decimal)
}
