package main

import (
	"github.com/spf13/cobra"
	"log"
)

var Root = &cobra.Command{
	Use:   "gok",
	Short: "gok is command line helper for go-kid projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Usage()
		}
		ctlName := args[0]

		return nil
	},
}

func main() {
	if err := Root.Execute(); err != nil {
		log.Fatal(err)
	}
}
