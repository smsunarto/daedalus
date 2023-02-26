package zk

import (
	"fmt"
	"os"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/smsunarto/daedalus/pkg/common"
)

func CompileCircuit(circuitName string, c common.CircuitEntry, buildDir string) error {
	fmt.Printf("Compiling circuit: %s\n", circuitName)

	// Compile circuit
	cs, err := frontend.Compile(c.Curve.ScalarField(), r1cs.NewBuilder, c.Circuit)
	if err != nil {
		return err
	}

	// Create output directory in ./build/<circuitName>
	outputDir := fmt.Sprintf("%s/%s", buildDir, circuitName)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	// Write the r1cs file in ./build/r1cs/<circuitName>.r1cs
	filepath := fmt.Sprintf("%s/%s.r1cs", outputDir, circuitName)
	WriteGnarkObjBinary(cs, filepath)

	return nil
}

func CompileAllCircuits(circuits map[string]common.CircuitEntry, buildDir string) error {
	for circuitName, circuit := range circuits {
		if err := CompileCircuit(circuitName, circuit, buildDir); err != nil {
			return err
		}
	}
	return nil
}
