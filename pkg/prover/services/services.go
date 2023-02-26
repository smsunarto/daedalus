package services

import (
	"context"
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/smsunarto/daedalus/pkg/prover/circuits"
)

type CircuitData struct {
	Circuit frontend.Circuit
	Curve   ecc.ID
	CS      constraint.ConstraintSystem
	PK      groth16.ProvingKey
	VK      groth16.VerifyingKey
}

type MockKVStore struct {
	CircuitData map[string]CircuitData
}

type IProver interface {
	Test(ctx context.Context) error
	GenerateProof(ctx context.Context, circuit frontend.Circuit, circuitName string) (groth16.Proof, error)
	VerifyProof(ctx context.Context, proof groth16.Proof, circuitName string, publicWitness witness.Witness) (bool, error)
}

type Prover struct{}

var (
	KVStore MockKVStore = MockKVStore{CircuitData: make(map[string]CircuitData)}
)

func NewService() Prover {
	var prover Prover
	prover.LoadCircuits()

	return Prover{}
}

func (p Prover) LoadCircuits() {
	keys := make([]string, 0, len(circuits.Circuits))
	for k := range circuits.Circuits {
		keys = append(keys, k)
	}

	for i := range keys {
		c := circuits.Circuits[keys[i]]
		fmt.Printf("Compiling circuit: %s\n", keys[i])
		r1cs, err := frontend.Compile(c.Curve.ScalarField(), r1cs.NewBuilder, c.Circuit)
		if err != nil {
			panic(err)
		}

		// generate trusted setup
		pk, vk, err := groth16.Setup(r1cs)
		if err != nil {
			panic(err)
		}

		KVStore.CircuitData[keys[i]] = CircuitData{c.Circuit, c.Curve, r1cs, pk, vk}
	}
}

func (p Prover) GenerateProof(fullWitness witness.Witness, circuitName string) (groth16.Proof, error) {
	proof, err := groth16.Prove(KVStore.CircuitData[circuitName].CS, KVStore.CircuitData["mimc"].PK, fullWitness)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func (Prover) VerifyProof(proof groth16.Proof, circuitName string, publicWitness witness.Witness) bool {
	err := groth16.Verify(proof, KVStore.CircuitData[circuitName].VK, publicWitness)
	return err == nil
}
