package zk

import (
	"bytes"
	"fmt"
	"os"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/smsunarto/daedalus/pkg/common"
)

func PerformTrustedSetup(circuitName string, c common.CircuitEntry, buildDir string) error {
	fmt.Printf("Performing trusted setup for circuit: %s\n", circuitName)

	// Create output directory in ./build/<circuitName>
	circuitDir := fmt.Sprintf("%s/%s", buildDir, circuitName)
	if err := os.MkdirAll(circuitDir, os.ModePerm); err != nil {
		return err
	}

	// Load constraint system from file
	fmt.Printf("Loading constraint system for: %s\n", circuitName)
	csFilepath := fmt.Sprintf("%s/%s.r1cs", circuitDir, circuitName)

	var buf bytes.Buffer
	err := LoadGnarkObjBinary(&buf, csFilepath)
	if err != nil {
		// Error: can't find r1cs file
		return err
	}

	cs := groth16.NewCS(c.Curve)
	cs.ReadFrom(&buf)

	// Generate trusted setup
	// TODO: trusted setup should be generated once using MPC, stored in a file
	// and loaded instead of being generated at runtime
	pk, vk, err := groth16.Setup(cs)
	if err != nil {
		return err
	}

	// Write the proving key and verify key file in ./build/<circuitName>/<circuitName>.(pk|vk)
	pkFilepath := fmt.Sprintf("%s/%s.pk", circuitDir, circuitName)
	vkFilepath := fmt.Sprintf("%s/%s.vk", circuitDir, circuitName)
	WriteGnarkObjBinary(pk, pkFilepath)
	WriteGnarkObjBinary(vk, vkFilepath)

	return nil
}

func PerformAllTrustedSetup(circuits map[string]common.CircuitEntry, buildDir string) error {
	for circuitName, circuit := range circuits {
		if err := PerformTrustedSetup(circuitName, circuit, buildDir); err != nil {
			return err
		}
	}

	return nil
}
