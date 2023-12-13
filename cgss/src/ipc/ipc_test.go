package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(request string, params string) *Response {
	response := new(Response)

	response.Code = "ECHO:" + request
	response.Body = params

	return response
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("From Client1", "client")
	resp2, _ := client2.Call("From Client2", "client")

	if resp1.Code != "ECHO:From Client1" || resp2.Code != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
