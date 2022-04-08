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
)

func ListeService() {
	var (
		out []byte
		err error
	)

	out, err = exec.Command("/bin/bash", "-c", "cat /proc/net/tcp | awk '{ print $2 }'").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
