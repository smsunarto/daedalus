package main

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/smsunarto/daedalus/pkg/common"
)

var Circuits map[string]common.CircuitEntry

func AddEntry(name string, circuit frontend.Circuit, curve ecc.ID) {
	if Circuits == nil {
		Circuits = make(map[string]common.CircuitEntry)
	}
	if _, ok := Circuits[name]; ok {
		panic("Circuit name " + name + " already taken by another deployed circuit")
	}

	Circuits[name] = common.CircuitEntry{Circuit: circuit, Curve: curve}
}

func main() {}
