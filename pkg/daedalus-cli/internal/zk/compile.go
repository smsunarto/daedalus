package zk

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/smsunarto/daedalus/pkg/daedalus-cli/circuits"
)

const (
	buildDir = "build"
)

func CompileCircuit(circuitName string, c circuits.CircuitEntry) error {
	fmt.Printf("Compiling circuit: %s\n", circuitName)

	cs, err := frontend.Compile(c.Curve.ScalarField(), r1cs.NewBuilder, c.Circuit)
	if err != nil {
		panic(err)
	}

	csBuf := new(bytes.Buffer)
	if _, err := cs.WriteTo(csBuf); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.r1cs", circuitName)
	filepath := fmt.Sprintf("%s/%s", buildDir, filename)

	if err := os.MkdirAll(buildDir, os.ModePerm); err != nil {
		return err
	}

	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	w := bufio.NewWriter(fo)
	if _, err := w.Write(csBuf.Bytes()); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func CompileAllCircuits(circuits map[string]circuits.CircuitEntry) error {
	for circuitName, circuit := range circuits {
		if err := CompileCircuit(circuitName, circuit); err != nil {
			return err
		}
	}
	return nil
}

// func PerformTrustedSetup() {
// 	keys := make([]string, 0, len(circuits.Circuits))
// 	for k := range circuits.Circuits {
// 		keys = append(keys, k)
// 	}

// 	// // Generate trusted setup
// 	// // TODO: trusted setup should be generated once using MPC, stored in a file
// 	// // and loaded instead of being generated at runtime
// 	// pk, vk, err := groth16.Setup(r1cs)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// KVStore.CircuitData[keys[i]] = CircuitData{c.Circuit, c.Curve, r1cs, pk, vk}
// }
