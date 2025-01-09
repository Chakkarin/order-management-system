package server

import (
	"fmt"
	"log"
)

func (s *ginServer) startHttp() {

	// s.app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// s.app.GET("/v1/health", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "OK")
	// })

	// Register HTTP handlers from all modules
	// for _, handler := range s.httpHandlers {
	// 	handler.HandlerHttp(s.app)
	// }

	port := fmt.Sprintf(`:%v`, s.cfg.ServicePort)

	log.Printf("🥳 HTTP server is running HTTP on port %s", port)
	if err := s.app.Run(port); err != nil {
		log.Panicf("❌ Failed to start HTTP server: %v", err)
	}

}
