package cmd

import (
	"github.com/smsunarto/daedalus/pkg/orchestrator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compileCmd)
}

var compileCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		orchestrator.Test()
	},
}
