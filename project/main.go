package main

import (
	"fmt"
	"driver"
    . "network"
    . "def"
)

func SEND(c chan Udp_message, msg Udp_message){
    c <- msg
    fmt.Println("Sending")
}

func RECEIVE(c chan Udp_message) {
    msg := <- c
    fmt.Println(msg.Data)
}


func main() {


    var localListenPort, broadcastListenPort, message_size int=20003,30000,1024
    send_ch := make(chan Udp_message)
    receive_ch := make(chan Udp_message)
    var msg=Udp_message{"broadcast","hei",1024}

    Udp_init(localListenPort, broadcastListenPort, message_size, send_ch, receive_ch)

    go RECEIVE(receive_ch)
    go SEND(send_ch, msg)
    

    // Initialize hardware
    if (driver.Elev_init() == 0) {
        fmt.Printf("Unable to initialize elevator hardware!\n")
    }

    fmt.Printf("Press STOP button to stop elevator and exit program.\n");

    driver.Elev_init()



   
    for {
            // Stop elevator and exit program if the stop button is pressed
        
        if (driver.Elev_get_stop_signal() == 1) {
            driver.Elev_set_motor_direction(driver.DIR_STOP)
            break

        }
    }
}

