package kit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"

	pb "github.com/happen-zhang/snippet/benckmark/grpc-gateway/proto"
)

type grpcServer struct {
	say grpctransport.Handler
}

func NewGRPCServer(ep endpoint.Endpoint) pb.HelloServer {
	return &grpcServer{
		say: grpctransport.NewServer(
			ep,
			decodeGRPCSayRequest,
			encodeGRPCSayReply,
		),
	}
}

func (srv *grpcServer) Say(ctx context.Context, req *pb.SayRequest) (*pb.SayReply, error) {
	_, resp, _ := srv.say.ServeGRPC(ctx, req)
	return resp.(*pb.SayReply), nil
}

func NewHTTPHandler(ep endpoint.Endpoint) http.Handler {
	http.Handle("/say", httptransport.NewServer(
		ep,
		decodeHTTPSayRequest,
		encodeHTTPSayReponse,
	))
	return http.DefaultServeMux
}

func decodeGRPCSayRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}

func encodeGRPCSayReply(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(string)
	return &pb.SayReply{
		Sentence: resp,
	}, nil
}

func decodeHTTPSayRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	// TODO convert
	return nil, nil
}

func encodeHTTPSayReponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Fprint(w, `{"sentence": "helloworld!"}`)
	return nil
}
