package main

import(
	"fmt"
	"net"
	"time"
	"encoding/json"
)

const (
	host = "129.241.187.255"
	udpPort = "20021"
	port_t = "30000"
)

type Message struct {
	N	int
	Str1	string
	Str2	[]string
}

func UDPReceive(port string) {
	buff := make([]byte, 1024)
	var mess_rec Message
	addr, _ := net.ResolveUDPAddr("udp", ":" + port)
	sock, _ := net.ListenUDP("udp", addr)
	for {
		_,_, err:=sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error UDP read: ", err)
		} 
		json.Unmarshal(buff, &mess_rec)
		
		fmt.Printf("%+v\n", mess_rec)
	}
}

func UDPSend(m []byte) {
	raddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, udpPort))
	if err != nil {
		fmt.Println("Failed to resolve address for: " + udpPort)
	}
	conn, err := net.DialUDP("udp", nil, raddr)	
	if err != nil {
		fmt.Println("Error dial: ", err)
	}
	go UDPReceive(udpPort)
	for {
		time.Sleep(1000*time.Millisecond)
		conn.Write(m)
		fmt.Println("Msg sent")	
	}
}

func main() {
	messages := Message{
		N:	1,
		Str1:	"string 2",
		Str2:	[]string{"s1","s2","s3"},	
	}
	b, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("error: ", err)
	}
	UDPSend(b)
	
	
	
	
	
}
