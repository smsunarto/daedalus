package tests

// func Test() error {
// 	KVStore = MockKVStore{make(map[string]CircuitData)}

// 	assignment := &circuits.MimcCircuit{
// 		PreImage: "16130099170765464552823636852555369511329944820189892919423002775646948828469",
// 		Hash:     "12886436712380113721405259596386800092738845035233065858332878701083870690753",
// 	}

// 	falseAssignment := &circuits.MimcCircuit{
// 		PreImage: "16130099170765464552823636852555369511329944820189892919423002775646948828469",
// 		Hash:     "123123",
// 	}

// 	fmt.Print("Loading circuits...\n")
// 	p.LoadCircuits(ctx)

// 	fmt.Print("Generating witness...\n")
// 	//// Generating witness for happy path
// 	// serialize full witness
// 	fullWitness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
// 	if err != nil {
// 		panic(err)
// 	}

// 	// extract public witness from full witness
// 	publicWitness, err := fullWitness.Public()
// 	if err != nil {
// 		panic(err)
// 	}

// 	//// Generating witness for unhappy path
// 	falseFullWitness, err := frontend.NewWitness(falseAssignment, ecc.BN254.ScalarField())
// 	if err != nil {
// 		panic(err)
// 	}

// 	falsePublicWitness, err := falseFullWitness.Public()
// 	if err != nil {
// 		panic(err)
// 	}
// 	////////////////////////////////////////////

// 	fmt.Print("Generating proof...\n")
// 	proof := p.GenerateProof(ctx, fullWitness, "mimc")

// 	fmt.Print("Verifying proof...\n")
// 	isVerified := p.VerifyProof(ctx, proof, "mimc", publicWitness)
// 	isVerifiedFalse := p.VerifyProof(ctx, proof, "mimc", falsePublicWitness)

// 	fmt.Printf("%t\n", isVerified)
// 	fmt.Printf("%t\n", isVerifiedFalse)

// 	return nil
// }
