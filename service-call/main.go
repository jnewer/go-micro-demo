package main

import (
	"context"
	"fmt"
	myhttp "github.com/go-micro/plugins/v4/client/http"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
)

func main() {
	register := consul.NewRegistry()
	s := selector.NewSelector(
		selector.Registry(register),
		selector.SetStrategy(selector.Random),
	)

	myClient := myhttp.NewClient(client.Selector(s), client.ContentType("application/json"))

	req := myClient.NewRequest("demo-http", "/demo", nil)
	resp := make(map[string]string)

	myClient.Call(context.Background(), req, &resp)
	fmt.Printf("resp: %v\n", resp)
}
