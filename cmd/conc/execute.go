package conc

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "concurency",
	Short: "Starts the backoffice reverse proxy",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	fmt.Println("Running from the concurency class")
}
