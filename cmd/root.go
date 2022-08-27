// cat /proc/net/tcp
// https://www.kernel.org/doc/Documentation/networking/proc_net_tcp.txt
// cat /proc/pid/comm
// lsof -i tcp | grep LISTEN | awk '{ print $2 }' # get pid
// hbci = :3000

package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/karchx/lsproc/internal/tui"
	"github.com/spf13/cobra"
)

var mode string

var rootCmd = &cobra.Command{
	Use:     "lsproc",
	Short:   "A terminal user interface for the open tcp process",
	Example: "lsproc",
	Run: func(cmd *cobra.Command, _ []string) {
    m := tui.New()
    p := tea.NewProgram(m)
    if err := p.Start(); err != nil {
      log.Fatal("Failed to start", err)
    }
		/*if err := tui.NewProgram().Start(); err != nil {
			fmt.Println("Could not start ui", err)
			os.Exit(1)
		}*/
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "", "Mode Execute")
}

func Execute() error {
	return rootCmd.Execute()
}
