package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	renderer "github.com/zackattackz/renderer/proto"
)

type Renderer struct{}

// Return a new handler
func New() *Renderer {
	return &Renderer{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Renderer) Call(ctx context.Context, req *renderer.Request, rsp *renderer.Response) error {
	log.Info("Received Renderer.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Renderer) Stream(ctx context.Context, req *renderer.StreamingRequest, stream renderer.Renderer_StreamStream) error {
	log.Infof("Received Renderer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&renderer.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Renderer) PingPong(ctx context.Context, stream renderer.Renderer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&renderer.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
