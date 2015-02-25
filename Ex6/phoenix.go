package main

import{
	"fmt"
	"net"
	"time"
}

const (
	host = "129.241.187.255"
	udpPort = "20021"
) 

/*func SpawnBackup() {
	Println("Spawning backup")
	cmd := exec.Command("mate-terminal", "-x", "go", "run", "phoenix.go")
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	print(string(out))
}*/

func Backup(conn *net.UDPConn) {
	buff := make([]byte, 256)
	
	for {
		select {
			case "dead":
				//new primary with correct countervalue
				//SpawnBackup()

			default:
				//save countervalue
		}
		
	}
}


func main() {
	var counter int = 0
	var data = make([]byte, 256)

	udpAddr, err = net.ResolveUDPAddr("udp", localhost)
	if err != nil {
		fmt.Println("Failed to resolve address")
	}

	conn, err := net.DialUDP("udp", nil, raddr)	
	if err != nil {
		fmt.Println("Error dial: ", err)
	}

	go Backup(conn)

	for {
		counter++
		fmt.Println(counter)
		time.Sleep(100*time.Millisecond)
		data[0] = byte(counter)
		conn.Write(data)
	}
	
}