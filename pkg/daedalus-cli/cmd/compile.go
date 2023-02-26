package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"plugin"

	"github.com/smsunarto/daedalus/pkg/common"
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
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			return err
		}

		// Make sure that this is the circuits root directory
		if _, err := os.Stat("daedalus.toml"); os.IsNotExist(err) {
			fmt.Println("daedalus.toml file does not exist in the current directory")
		} else {
			fmt.Println("daedalus.toml file exists in the current directory")
		}

		// Create a build directory in the circuits root directory
		buildDir := fmt.Sprintf("%s/build", wd)
		if err := os.MkdirAll(buildDir, os.ModePerm); err != nil {
			return err
		}

		// Compile the circuits into a shared object binary
		soOutputPath := fmt.Sprintf("%s/circuits.so", buildDir)
		goBuildCmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soOutputPath)
		out, err := goBuildCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(out))

		// Load the shared object binary
		p, err := plugin.Open(soOutputPath)
		if err != nil {
			panic(err)
		}

		// Serialize the Circuits map
		c, err := p.Lookup("Circuits")
		if err != nil {
			panic(err)
		}
		circuits := *c.(*map[string]common.CircuitEntry)

		// Compile the circuits into r1cs files
		if err := zk.CompileAllCircuits(circuits, buildDir); err != nil {
			return err
		}

		return nil
	},
}
