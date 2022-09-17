package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"service-pb/handler"
	pb "service-pb/proto"
)

func main() {
	service := micro.NewService(micro.Name("sum-service"))
	pb.RegisterSumServiceHandler(service.Server(), handler.NewHandler())

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}
