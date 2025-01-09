package server

import (
	"fmt"
	"log"
	"net"
)

func (s *ginServer) startGrpc() {

	port := fmt.Sprintf(`:%v`, s.cfg.GrpcPort)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("❌ Failed to listen on port %s: %v", port, err)
	}

	log.Printf("🥳 gRPC server is running gRPC on port %s", port)
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Printf("❌ Failed to serve gRPC server: %v", err)
	}

}
