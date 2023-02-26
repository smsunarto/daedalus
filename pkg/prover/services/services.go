package services

import (
	"context"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
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

// TODO: this should be replaced with a real persistence layer
var (
	KVStore MockKVStore = MockKVStore{CircuitData: make(map[string]CircuitData)}
)

func NewService() Prover {
	return Prover{}
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
