package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}

	b, err := json.Marshal(req)
	if err != nil {
		return
	}

	// 和Server进行沟通
	client.conn <- string(b)
	str := <-client.conn // 等待返回值

	err = json.Unmarshal([]byte(str), &resp)
	if err != nil {
		fmt.Println("Unmarshal failed!")
	}

	return resp, err
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
