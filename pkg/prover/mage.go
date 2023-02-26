//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Start up all the processes in development mode.
func Run() {
	mg.Deps(Build)
	fmt.Println("Running...")
	sh.RunV("./prover")
	os.Remove("./prover")
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "prover", ".")
	return cmd.Run()
}

// // A custom install step if you need your bin someplace other than go/bin
// func Install() error {
// 	mg.Deps(Build)
// 	fmt.Println("Installing...")
// 	return os.Rename("./orchestrator", "/usr/bin/orchestrator")
// }

// // Clean up after yourself
// func Clean() {
// 	fmt.Println("Cleaning...")
// 	os.RemoveAll("orchestrator")
// }
