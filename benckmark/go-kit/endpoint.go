package kit

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type SayRequest struct {
	Word   string
	Repeat int32
}

func MakeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*SayRequest)
		return svc.Say(ctx, req.Word, int(req.Repeat)), nil
	}
}
