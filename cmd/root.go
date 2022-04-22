// cat /proc/net/tcp
// https://www.kernel.org/doc/Documentation/networking/proc_net_tcp.txt
// cat /proc/pid/comm
// lsof -i tcp | grep LISTEN | awk '{ print $2 }' # get pid
// hbci = :3000

// TODO: add cobra to handle modes: cmd and ui

package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	mode    string
	rootCmd = &cobra.Command{
		Use:   "lsproc",
		Short: "Use mode ui or cli",
	}
	host    = "127.0.0.1"
	ports   = "1-65535"
	threads = 1000
	timeout = 1 * time.Second
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "", "Mode Execute")
}

func Execute() error {
	return rootCmd.Execute()
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
