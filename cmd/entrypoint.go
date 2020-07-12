package main

import (
	"github.com/abatilo/cobra-no-init/cmd/simple"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "entrypoint",
		Short: "This is a demonstration for how I register my cobra applications",
	}

	root.AddCommand(simple.Cmd())
	root.Execute()
}
