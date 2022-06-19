package main

import (
	"github.com/zackattackz/renderer/handler"
	pb "github.com/zackattackz/renderer/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("renderer"),
	)

	// Register handler
	pb.RegisterRendererHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
