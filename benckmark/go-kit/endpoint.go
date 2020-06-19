package kit

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return svc.Say(), nil
	}
}
