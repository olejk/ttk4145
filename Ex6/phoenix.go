package main

import{
	"fmt"
	"net"
	"time"
	"os/exec"
}

const (
	host = "129.241.187.255"
	port = "20013"
)

func SpawnProcess() {
	Println("Spawning backup")
	cmd := exec.Command("gnome-terminal", "-x", "go", "run", "phoenix.go")
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	print(string(out))
}

func slave(sock *net.UDPConn, masterAlive bool, counter *int) bool{
	for(masterAlive){
		sock.SetReadDeadline(time.Now().Add(2*time.Second))
		data := make([]byte, 256)
		n, _, err := sock.ReadFromUDP(data[0:])
		if err != nil {
			masterAlive = false
			return masterAlive
		} else {
			*count = getCount(string(data[:n]))
			fmt.Println("Slave, master count:", *count)
		}
	}
	return true
}

func main() {
	masterAlive := true
	counter := 0
	t_count := 0
	udpAddr, _ = net.ResolveUDPAddr("udp", ":" + port)
	sock, _ := net.ListenUDP("udp", udpAddr)

	masterAlive = slave(sock, masterAlive, &counter)
	sock.Close()

	SpawnProcess()
	t_count = counter
	addr, _ := net.ResolveUDPAddr("udp4", host + ":" + port)
	sock2, _ := net.DialUDP("udp4", nil, addr)

	for {
		msg := "Count:" + strconv.Itoa(counter)
		_, err := mConn.Write([]byte(msg))
		fmt.Println("Master count: ", counter)
		if err != nil {
			fmt.Println("Error:Broadcast", err.Error())
		}
		counter++
		time.Sleep(time.Second)
	}
	mConn.Close()
}