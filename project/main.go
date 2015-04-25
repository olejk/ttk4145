package main

import (
	"fmt"
	. "driver"
    . "network"
    . "def"
    . "eventDetection"
    . "timer"
)

func SEND(c chan Udp_message, msg Udp_message){
    c <- msg
    fmt.Println("Sending")
}

func RECEIVE(c chan Udp_message) {
    msg := <- c
    fmt.Println(msg.Data)
}

func initStateMachine(){
    Msg.State = IDLE
    Msg.Dir = DIR_UP
    Msg.PrevFloor = Elev_get_floor_sensor_signal()
}

func main() {
    timerChan := make(chan string)
    timeOutChan := make(chan int)
    doneChan := make(chan string)
    
   
    // var localListenPort, broadcastListenPort, message_size int=20003,30000,1024
    // send_ch := make(chan Udp_message)
    // receive_ch := make(chan Udp_message)
    // var msg=Udp_message{"broadcast","hei",1024}

    // Udp_init(localListenPort, broadcastListenPort, message_size, send_ch, receive_ch)

    // go RECEIVE(receive_ch)
    // go SEND(send_ch, msg)
    

    // Initialize hardware
    if (Elev_init() == 0) {
        fmt.Printf("Unable to initialize elevator hardware!\n")
    }
    initStateMachine()

    go DoorTimer(timerChan, timeOutChan)
    go EventHandler(timerChan, timeOutChan)
    
    fmt.Println(<-doneChan)
}

