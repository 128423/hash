package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/128423/hash/server/database"
	"github.com/128423/hash/server/pb"
)

type Server struct {
}

func (s *Server) GetDiscount(ctx context.Context, req *pb.RequestDiscount) (*pb.ResponseDiscount, error) {
	User, err := database.GetUser(req.GetUserId())
	if err != nil {
		return nil, errors.New("User not found")
	}
	product, err := database.GetProduct(req.GetProductId())
	if err != nil {
		return nil, errors.New("Product not found")
	}
	t := time.Now()

	if t.Day() == 25 && t.Month() == time.November {
		product.Discount.Ptc = 10.0
		product.Discount.ValueInCents = int64(float64(product.PriceCents) * 0.10)
	} else if User.DateBirth.Day() == t.Day() && User.DateBirth.Month() == t.Month() {
		product.Discount.Ptc = 5.0
		product.Discount.ValueInCents = int64(float64(product.PriceCents) * 0.05)
	} else {
		product.Discount.Ptc = 3.0
		product.Discount.ValueInCents = int64(float64(product.PriceCents) * 0.03)
	}

	return &pb.ResponseDiscount{
		Discount: &pb.Discount{
			Pct:          product.Discount.Ptc,
			ValueInCents: product.Discount.ValueInCents,
		},
	}, nil
}
