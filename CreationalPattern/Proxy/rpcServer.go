package proxy

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// 启动数据库，对外提供RPC接口进行数据库的访问
func Start() {
	rpcServer := rpc.NewServer()
	server := &Server{data: make(map[string]string)}
	if err := rpcServer.Register(server); err != nil {
		fmt.Printf("Register Server to rpc failed,error: %v", err)
		return
	}
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Listen tcp failed, error: %v", err)
		return
	}
	go rpcServer.Accept(l)
	time.Sleep(1 * time.Second)
	fmt.Println("Rpc server start success.")
}
