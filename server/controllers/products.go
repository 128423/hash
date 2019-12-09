package controllers

import (
	"context"

	"github.com/128423/hash/server/pb"
)

type Server struct {
}

func (s *Server) GetDiscount(ctx context.Context, req *pb.RequestDiscount) (*pb.ResponseDiscount, error) {
	return nil, nil
}
