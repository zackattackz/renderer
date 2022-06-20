package handler

import (
	"context"

	"github.com/micro/micro/v3/service/broker"
	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/proto"

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

func (e *Renderer) BrokerHandler(ctx context.Context) broker.Handler {
	return func(m *broker.Message) error {
		var req renderer.Request
		var rsp renderer.Response
		err := proto.Unmarshal(m.Body, &req)
		if err != nil {
			return err
		}

		err = e.Call(ctx, &req, &rsp)
		if err != nil {
			return err
		}

		rspProto, err := proto.Marshal(&rsp)
		if err != nil {
			return err
		}

		return broker.Publish("renderer.Response", &broker.Message{Body: rspProto})
	}
}
