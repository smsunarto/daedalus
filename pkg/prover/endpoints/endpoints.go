package endpoints

import (
	"context"
	"errors"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/go-kit/kit/endpoint"
	"github.com/smsunarto/daedalus/pkg/prover/services"
)

type Endpoints struct {
	GenerateProofEndpoint endpoint.Endpoint
}

func NewEndpoints(service services.Prover) Endpoints {
	return Endpoints{
		GenerateProofEndpoint: MakeGenerateProofEndpoint(service),
	}
}

func MakeGenerateProofEndpoint(prover services.Prover) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func() {
			if v := recover(); v != nil {
				response = GenerateProofResponse{Proof: nil, Err: errors.New(v.(string))}
				err = errors.New(v.(string))
			}
		}()

		req := request.(GenerateProofRequest)
		proof, err := prover.GenerateProof(req.FullWitness, req.CircuitName)
		return GenerateProofResponse{Proof: proof, Err: err}, nil
	}
}

func (ep Endpoints) GenerateProof(ctx context.Context, fullWitness witness.Witness, circuitName string) (groth16.Proof, error) {
	resp, err := ep.GenerateProofEndpoint(ctx, GenerateProofRequest{FullWitness: fullWitness, CircuitName: circuitName})
	if err != nil {
		return nil, err
	}
	response := resp.(GenerateProofResponse)
	return response.Proof, response.Err
}

type GenerateProofRequest struct {
	FullWitness witness.Witness
	CircuitName string
}

type GenerateProofResponse struct {
	Proof groth16.Proof `json:"v"`
	Err   error         `json:"-"`
}
