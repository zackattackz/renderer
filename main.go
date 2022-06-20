package main

import (
	"context"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/broker"
	"github.com/micro/micro/v3/service/logger"
	"github.com/zackattackz/renderer/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("statikit.svc.renderer"),
	)

	// Register handler
	//pb.RegisterRendererHandler(srv.Server(), handler.New())

	broker.Subscribe("renderer.Request", handler.New().BrokerHandler(context.Background()))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
