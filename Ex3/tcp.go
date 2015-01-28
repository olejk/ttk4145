package main

import(
	"fmt"
	"net"
	//"os"
	//"bufio"
)

const (
	server = "129.241.187.153"
	port_fixed = "34933"
	port_delim = "33546"
)

func connTCP() { // Port-parameter for fixed/delim
	conn, err := net.Dial("tcp", net.JoinHostPort(server, port_fixed))
	if err != nil {
		fmt.Println("Error connecting to TCP server")
	}
	
	addr, err := net.ResolveTPCAddr("tcp", server)
	if err != nil {
		fmt.Println("Failed to resolve address for: " + port_fixed)
	}

	listener, err := net.ListenTCP("tcp", addr)
}

func main() {
	
	
}
