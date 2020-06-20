package kit

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

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

func decodeGRPCSayRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SayRequest)
	return &SayRequest{
		Word:   req.Word,
		Repeat: req.RepeatCount,
	}, nil
}

func encodeGRPCSayReply(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(string)
	return &pb.SayReply{
		Sentence: resp,
	}, nil
}

func decodeHTTPSayRequest(_ context.Context, request *http.Request) (interface{}, error) {
	query := request.URL.Query()
	repeat, _ := strconv.Atoi(query.Get("repeatCount"))
	return &SayRequest{
		Word:   query.Get("word"),
		Repeat: int32(repeat),
	}, nil
}

func encodeHTTPSayReponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Fprint(w, `{"sentence": "hello, `+response.(string)+`!"}`)
	return nil
}
