package handler

import (
	"context"
	pb "service-pb/proto"
	"service-pb/service"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) Sum(context context.Context, in *pb.Request, out *pb.Response) error {
	inputs := make([]int64, 0)
	var i int64 = 0

	for ; i < in.Num; i++ {
		inputs = append(inputs, i)
	}

	out.Num = service.GetSum(inputs...)
	return nil
}
