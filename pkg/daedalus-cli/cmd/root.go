package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daedalus",
	Short: "Daedalus CLI is a command line interface for Daedalus",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Daedalus CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
