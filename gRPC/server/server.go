package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"high-traffic-practice/config"
	"high-traffic-practice/gRPC/paseto"
	auth "high-traffic-practice/gRPC/proto"
	"log"
	"net"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if listen, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {
		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})

		reflection.Register(server)

		go func() {
			log.Println("gRPC server is running on", cfg.GRPC.URL)
			if err := server.Serve(listen); err != nil {
				panic(err)
			}
		}()
	}

	return nil
}

func (s *GRPCServer) CreateAuth(_ context.Context, req *auth.CreateTokenRequest) (*auth.CreateTokenResponse, error) {
	data := req.Auth
	token := data.Token
	s.tokenVerifyMap[token] = data

	return &auth.CreateTokenResponse{
		Auth: data,
	}, nil
}

func (s *GRPCServer) VerifyAuth(_ context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	token := req.Token

	res := &auth.VerifyTokenResponse{
		Verify: &auth.Verify{
			Auth: nil,
		}}

	if authData, ok := s.tokenVerifyMap[token]; !ok {
		res.Verify.Status = auth.ResponseType_FAILED
	} else if authData.ExpireDate < authData.CreateDate {
		delete(s.tokenVerifyMap, token)
		res.Verify.Status = auth.ResponseType_EXPIRED
	} else {
		res.Verify.Status = auth.ResponseType_SUCCESS
	}
	return res, nil
}
