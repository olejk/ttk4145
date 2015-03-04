package main

import (
	"fmt"
	"net"
	"time"
	//"os/exec"
	"strconv"
	"strings"
)

const (
	host = "129.241.187.255"
	port = "20013"
)

/*func SpawnProcess() {
	fmt.Println("Spawning backup")
	cmd := exec.Command("gnome-terminal", "-x", "go", "run", "phoenix.go")
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	print(string(out))
}*/

func slave(sock *net.UDPConn, masterAlive bool, counter *int) bool{
	for(masterAlive){
		sock.SetReadDeadline(time.Now().Add(2*time.Second))
		data := make([]byte, 256)
		n, _, err := sock.ReadFromUDP(data[0:])
		if err != nil {
			masterAlive = false
			return masterAlive
		} else {
			s := strings.TrimLeft(string(data[:n]), "Count: ")
			count,_ := strconv.Atoi(s)
			*counter = count
			fmt.Println("Slave, master count:", *counter)
		}
	}
	return true
}

func main() {
	masterAlive := true
	counter := 0
	t_count := 0
	udpAddr, _ := net.ResolveUDPAddr("udp", host + ":" + port)
	sock, _ := net.ListenUDP("udp", udpAddr)
	masterAlive = slave(sock, masterAlive, &counter)
	sock.Close()

	//SpawnProcess()
	t_count = counter
	fmt.Println(t_count)
	addr, _ := net.ResolveUDPAddr("udp4", host + ":" + port)
	sock2, _ := net.DialUDP("udp4", nil, addr)

	for {
		msg := "Count:" + strconv.Itoa(counter)
		_, err := sock2.Write([]byte(msg))
		fmt.Println("Master count: ", counter)
		if err != nil {
			fmt.Println("Error:Broadcast", err.Error())
		}
		counter++
		time.Sleep(time.Second)
	}
	sock2.Close()
}
