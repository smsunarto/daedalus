package orchestrator

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/smsunarto/daedalus/circuits"
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

var KVStore MockKVStore

func LoadCircuits() {

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

func GenerateProof(fullWitness witness.Witness, circuitName string) groth16.Proof {
	proof, err := groth16.Prove(KVStore.CircuitData[circuitName].CS, KVStore.CircuitData["mimc"].PK, fullWitness)
	if err != nil {
		panic(err)
	}

	fmt.Printf("proof: %v)", proof)
	return proof
}

func VerifyProof(proof groth16.Proof, circuitName string, publicWitness witness.Witness) bool {
	err := groth16.Verify(proof, KVStore.CircuitData[circuitName].VK, publicWitness)
	return err == nil
}

func Test() {
	KVStore = MockKVStore{make(map[string]CircuitData)}

	assignment := &circuits.MimcCircuit{
		PreImage: "16130099170765464552823636852555369511329944820189892919423002775646948828469",
		Hash:     "12886436712380113721405259596386800092738845035233065858332878701083870690753",
	}

	falseAssignment := &circuits.MimcCircuit{
		PreImage: "16130099170765464552823636852555369511329944820189892919423002775646948828469",
		Hash:     "123123",
	}

	fmt.Print("Loading circuits...\n")
	LoadCircuits()

	fmt.Print("Generating witness...\n")
	//// Generating witness for happy path
	// serialize full witness
	fullWitness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}

	// extract public witness from full witness
	publicWitness, err := fullWitness.Public()
	if err != nil {
		panic(err)
	}

	//// Generating witness for unhappy path
	falseFullWitness, err := frontend.NewWitness(falseAssignment, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}

	falsePublicWitness, err := falseFullWitness.Public()
	if err != nil {
		panic(err)
	}
	////////////////////////////////////////////

	fmt.Print("Generating proof...\n")
	proof := GenerateProof(fullWitness, "mimc")

	fmt.Print("Verifying proof...\n")
	isVerified := VerifyProof(proof, "mimc", publicWitness)
	isVerifiedFalse := VerifyProof(proof, "mimc", falsePublicWitness)

	fmt.Printf("%t\n", isVerified)
	fmt.Printf("%t\n", isVerifiedFalse)
}
