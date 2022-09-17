package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/selector"
)

func main() {
	registry := consul.NewRegistry()
	services, _ := registry.GetService("demo-http")
	next := selector.Random(services)
	node, _ := next()
	fmt.Printf("node.Address: %v\n", node.Address)
	fmt.Printf("node.Id: %v\n", node.Id)
}
