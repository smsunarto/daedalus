package circuits

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

type CircuitEntry struct {
	Circuit frontend.Circuit
	Curve   ecc.ID
}

var Circuits map[string]CircuitEntry

func addEntry(name string, circuit frontend.Circuit, curve ecc.ID) {

	if Circuits == nil {
		Circuits = make(map[string]CircuitEntry)
	}
	if _, ok := Circuits[name]; ok {
		panic("name " + name + "already taken by another deployed circuit ")
	}

	Circuits[name] = CircuitEntry{circuit, curve}
}
