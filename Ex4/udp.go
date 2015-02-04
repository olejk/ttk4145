package main

import(
	"fmt"
	"net"
	"time"
	"encoding/json"
	"os"
)


const (
	host = "129.241.187.255"
	udpPort = "20021"
	port_t = "30000"
)

type Message struct {
	str1	string
	str2	string
	str3	string
}

func udpReceive(port string, m []byte) {
	buff := make([]byte, 1024)
	var mess_rec []Messages
	addr, _ := net.ResolveUDPAddr("udp", ":" + port)
	sock, _ := net.ListenUDP("udp", addr)
	for {
		_,_, err:=sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error UDP read: ", err)
		} 
		err := json.Unmarshal(m, &mess_rec)
		fmt.Println("%+v", mess_rec)
	}
}

func udpSend(m []byte) {
	raddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, udpPort))
	if err != nil {
		fmt.Println("Failed to resolve address for: " + udpPort)
	}
	conn, err := net.DialUDP("udp", nil, raddr)	
	if err != nil {
		fmt.Println("Error dial: ", err)
	}
	
	go udpReceive(udpPort)
	for {
		time.Sleep(1000*time.Millisecond)
		conn.Write(m)
		fmt.Println("Msg sent")	
	}

}

func main() {
	messages := Message{
		str1:	"string 1",
		str2:	"string 2",
		str3:	"string 3",	
	}
	b, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("error: ", err)
	}
	udpSend(messages)
}
