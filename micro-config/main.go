package main

import (
	"fmt"
	yaml "github.com/go-micro/plugins/v4/config/encoder/yaml"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
)

func main() {
	//if err := config.Load(file.NewSource(file.WithPath("./config.json"))); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	enc := yaml.NewEncoder()
	c, _ := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(enc))))
	if err := c.Load(file.NewSource(file.WithPath("./config.yaml"))); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data", c.Map())

	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host

	if err := c.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
}
