package cmd

import (
	"fmt"
	"os"

	"github.com/samoei/toolbox/cmd/concurency"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the program",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(concurency.ConcurencyCommand)
	rootCmd.AddCommand(concurency.UserProfileCommand)
}
