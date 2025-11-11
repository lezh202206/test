package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "http://1.12.247.124:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("无法连接 Consul: %v", err)
	}

	// 注册服务
	registerService(client)

	// 服务发现
	discoverService(client, "example-service")

	// KV 操作
	keyValueOperations(client)
}

func registerService(client *api.Client) {
	registration := &api.AgentServiceRegistration{
		ID:      "example-go-service",
		Name:    "example-service",
		Address: "1.12.247.124",
		Port:    8080,
		//Check: &api.AgentServiceCheck{
		//	HTTP:     "http://127.0.0.1:8080/health",
		//	Interval: "10s",
		//	Timeout:  "1s",
		//},
	}
	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("服务注册失败: %v", err)
	}
	fmt.Println("服务注册成功！")
}

func discoverService(client *api.Client, serviceName string) {
	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		log.Fatalf("服务发现失败: %v", err)
	}
	for _, service := range services {
		fmt.Printf("发现服务: %s (%s:%d)\n",
			service.Service.Service,
			service.Service.Address,
			service.Service.Port)
	}
}

func keyValueOperations(client *api.Client) {
	kv := client.KV()
	p := &api.KVPair{Key: "example/key", Value: []byte("Hello, Consul!")}
	_, err := kv.Put(p, nil)
	if err != nil {
		log.Fatalf("写入 KV 失败: %v", err)
	}
	fmt.Println("KV 写入成功")
	pair, _, err := kv.Get("example/key", nil)
	if err != nil {
		log.Fatalf("读取 KV 失败: %v", err)
	}
	fmt.Printf("KV 读取结果: %s\n", pair.Value)
}
