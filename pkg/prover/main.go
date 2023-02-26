package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	stdlog "log"

	kitlog "github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/smsunarto/daedalus/pkg/prover/endpoints"
	"github.com/smsunarto/daedalus/pkg/prover/proto"
	"github.com/smsunarto/daedalus/pkg/prover/services"
	"github.com/smsunarto/daedalus/pkg/prover/transports"
	"google.golang.org/grpc"
)

func main() {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))

	// fs := flag.NewFlagSet("prover", flag.ExitOnError)
	// var (
	// 	grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
	// )

	var (
		service    = services.NewService()
		endpoints  = endpoints.NewEndpoints(service)
		grpcServer = transports.NewGRPCServer(endpoints)
	)

	var g group.Group
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", ":8082")
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", ":8082")
			baseServer := grpc.NewServer()
			proto.RegisterProverServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())

}
