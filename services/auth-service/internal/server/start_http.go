package server

import (
	"fmt"
	"log"
)

func (s *ginServer) startHttp() {

	port := fmt.Sprintf(`:%v`, s.cfg.ServicePort)

	log.Printf("🥳 HTTP server is running HTTP on port %s", port)
	if err := s.app.Run(port); err != nil {
		log.Panicf("❌ Failed to start HTTP server: %v", err)
	}

}
