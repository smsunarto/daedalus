package transports

import (
	"bytes"
	"context"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/witness"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/smsunarto/daedalus/pkg/prover/endpoints"
	"github.com/smsunarto/daedalus/pkg/prover/proto"
)

type GRPCServer struct {
	generateProof gt.Handler
	proto.UnimplementedProverServer
}

func NewGRPCServer(endpoints endpoints.Endpoints) proto.ProverServer {
	return &GRPCServer{
		generateProof: gt.NewServer(
			endpoints.GenerateProofEndpoint,
			decodeGenerateProofRequest,
			encodeGenerateProofResponse,
		),
	}
}

func (s *GRPCServer) GenerateProof(ctx context.Context, req *proto.GenerateProofRequest) (*proto.GenerateProofResponse, error) {
	_, resp, err := s.generateProof.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.GenerateProofResponse), nil
}

func decodeGenerateProofRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GenerateProofRequest)
	buf := bytes.NewBuffer(req.FullWitness)

	fullWitness, _ := witness.New(ecc.BN254.ScalarField())
	fullWitness.ReadFrom(buf)
	return endpoints.GenerateProofRequest{FullWitness: fullWitness, CircuitName: req.CircuitName}, nil
}

func encodeGenerateProofResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GenerateProofResponse)

	var resBuf bytes.Buffer
	resp.Proof.WriteTo(&resBuf)

	return &proto.GenerateProofResponse{Proof: resBuf.Bytes(), Err: resp.Err.Error()}, nil
}
