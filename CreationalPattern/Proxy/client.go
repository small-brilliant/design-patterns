package proxy

import (
	"fmt"
	"net/rpc"
)

type Client struct {
	cli *rpc.Client
}

func (c *Client) Save(record Record, reply *bool) error {
	var ret bool
	// 通过RPC调用服务端的接口
	err := c.cli.Call("Server.Save", record, &ret)
	if err != nil {
		fmt.Printf("Call db Server.Save rpc failed, error: %v", err)
		*reply = false
		return err
	}
	*reply = ret
	return nil
}

func (c *Client) Get(key string, reply *string) error {
	var ret string
	// 通过RPC调用服务端的接口
	err := c.cli.Call("Server.Get", key, &ret)
	if err != nil {
		fmt.Printf("Call db Server.Get rpc failed, error: %v", err)
		*reply = ""
		return err
	}
	*reply = ret
	return nil
}

// 工厂方法，返回远程代理实例
func CreateClient() *Client {
	rpcCli, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Create rpc client failed, error: %v.", err)
		return nil
	}
	return &Client{cli: rpcCli}
}
