package kv

import (
	"algs4/config"
	"log"
	"net"
	"net/rpc"
	"sync"
)

const (
	serviceName = "KVService"

	OK       = "OK"
	NotFound = "NotFound"
)

type PutRequest struct {
	Key   string
	Value string
}

type PutReply struct {
	Err string
}

type GetRequest struct {
	Key string
}

type GetReply struct {
	Value string
	Err   string
}

type ServiceSignature interface {
	Put(PutRequest, *PutReply) error
	Get(GetRequest, *GetReply) error
}

func registerService(svc ServiceSignature) error {
	return rpc.RegisterName(serviceName, svc)
}

func RunServer() {
	err := registerService(&Handler{
		mu: sync.Mutex{},
		kv: make(map[string]string),
	})
	if err != nil {
		log.Fatal("register err:", err)
	}
	listener, err := net.Listen(config.TCP, config.DefaultAddr)
	if err != nil {
		log.Fatal("listen err:", err)
	}
	defer listener.Close()

	log.Println("kv service started at", config.DefaultAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
