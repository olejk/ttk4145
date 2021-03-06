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
	myIp = "78.91.46.142"
)

type Message struct {
	N 		int
	Str1	string
	Str2	[]string
}

func UDPReceive(port string) {
	buff := make([]byte, 1024)
	var mess_rec Message
	addr, _ := net.ResolveUDPAddr("udp", ":" + port)
	sock, _ := net.ListenUDP("udp", addr)
	for {
		n,_, err:=sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error UDP read: ", err)
		} 
		json.Unmarshal(buff[:n], &mess_rec)
		fmt.Printf("Received: %+v\n", mess_rec)
	}
}

func UDPSend(m []byte, IP string) {
	raddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(IP, udpPort))
	if err != nil {
		fmt.Println("Failed to resolve address for: " + udpPort)
	}
	conn, err := net.DialUDP("udp", nil, raddr)	
	if err != nil {
		fmt.Println("Error dial: ", err)
	}
	for {
		time.Sleep(1000*time.Millisecond)
		conn.Write(m)
		fmt.Println("Msg sent")	
	}
}

func main() {
	messages := Message{
		N:		1,
		Str1:	"string 2",
		Str2:	[]string{"s1","s2","s3"},	
	}
	b, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("error: ", err)
	}

	go UDPReceive(udpPort)
	UDPSend(b, myIp)
}
