package comm

import (
	"fmt"
	"log"
	"net"
	"streamax-go/controller"
	"streamax-go/interfaces"
	"time"
)

//TCPServer struct
type TCPServer struct {
	Host                   string
	Port                   int
	onNewConnectionHandler func(c *Client)
	Listener               net.Listener
}

func (server *TCPServer) Listen() {

	defer server.Listener.Close()

	for {
		conn, _ := server.Listener.Accept()
		client := &Client{
			Connection:  conn,
			ConnectedAt: time.Now().UTC(),
		}
		//server.onNewConnectionHandler(client)
		go client.Listen()
	}
}

//NewTCPServer creates new instance of tcp server
func NewTCPServer(host string, port int) *TCPServer {
	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	log.Println("Creating server with address ", host, ":", port, "Error:", err)

	server := &TCPServer{
		Host:     host,
		Port:     port,
		Listener: l,
	}
	ServerCounters.AddFloat("Transmitted", 0)
	ServerCounters.AddFloat("Received", 0)
	DeviceChannelMap = make(map[string]*interfaces.IChannel)
	server.setOnNewClient(func(c *Client) {})
	controller.InitModuleMap()
	return server
}

//setOnNewClient indicates new client connected
func (server *TCPServer) setOnNewClient(callback func(c *Client)) {
	server.onNewConnectionHandler = callback
}
