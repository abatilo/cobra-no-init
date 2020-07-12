package simple

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd is the wrapped command for this package
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simple",
		Short: "This example has no dependencies",
		Run: func(_ *cobra.Command, _ []string) {
			run()
		},
	}
	return cmd
}

func run() {
	fmt.Println("Execution of a very simple command")
}
