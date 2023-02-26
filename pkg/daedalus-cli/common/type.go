package common

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

type CircuitEntry struct {
	Circuit frontend.Circuit
	Curve   ecc.ID
}
