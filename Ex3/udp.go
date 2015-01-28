package main

import(
	"fmt"
	"net"
)


const (
	port = ":30000"
)

func main() {
	buff := make([]byte, 1024)
	addr, _ := net.ResolveUDPAddr("udp", port)
	sock, _ := net.ListenUDP("udp", addr)
	for {
		_,_, err:=sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		} 
		fmt.Println(string(buff[:]))
	}
}
