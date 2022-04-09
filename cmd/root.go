// cat /proc/net/tcp
// https://www.kernel.org/doc/Documentation/networking/proc_net_tcp.txt
// cat /proc/pid/comm
// lsof -i tcp | grep LISTEN | awk '{ print $2 }' # get pid
// hbci = :3000

// TODO: add cobra to handle modes: cmd and ui

package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	mode    string
	rootCmd = &cobra.Command{
		Use:   "lsproc",
		Short: "Use mode ui or cli",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "", "Mode Execute")
}

func Execute() error {
	return rootCmd.Execute()
}

func decryptPort(hex string) int32 {
	decimal, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int32(decimal)
}

func ListenService() {
	var (
		out []byte
		err error
	)

	out, err = exec.Command("/bin/bash", "-c", "cat /proc/net/tcp | awk '{ print $2 }'").Output()
	if err != nil {
		log.Fatal(err)
	}

	pidHex := strings.Split(string(out), ":")

	for _, value := range pidHex {
		// port := decryptPort(value)
		fmt.Println(value)
	}
}
