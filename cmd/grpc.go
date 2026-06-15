package cmd

import (
	"e-wallet-transaction-service/helpers"
	"log"
	"net"

	"google.golang.org/grpc"
)

func ServerGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to start gRPC server:", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to start gRPC server:", err)
	}
}
