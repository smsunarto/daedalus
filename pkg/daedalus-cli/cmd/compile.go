package cmd

import (
	"github.com/smsunarto/daedalus/pkg/daedalus-cli/circuits"
	"github.com/smsunarto/daedalus/pkg/daedalus-cli/internal/zk"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compileCmd)
}

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile all circuits",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := zk.CompileAllCircuits(circuits.Circuits); err != nil {
			return err
		}
		return nil
	},
}
