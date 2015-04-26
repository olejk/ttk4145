package eventDetection

import (
	"fmt"
	"time"
	. "def"
	. "driver"
	. "stateMachine"
	. "queue"
	. "encdec"
	. "network"
	//"net"
)

func buttonEventDetector(orderEventChannel chan Order) {
	var currentSignalMatrix 	[3][N_FLOORS]int 
	var previousSignalMatrix 	[3][N_FLOORS]int 
	
	for {
		for floor:=0;floor<N_FLOORS;floor++ {
			for button:=0;button<3;button++ {
				currentSignalMatrix[button][floor] = Elev_get_button_signal(button,floor)
				if (currentSignalMatrix[button][floor] == 1 && previousSignalMatrix[button][floor] == 0) {
					orderEventChannel <- Order{floor, button}
				}
				previousSignalMatrix[button][floor] = currentSignalMatrix[button][floor]
			}
		}
		time.Sleep(10*time.Millisecond)
	}
}


func floorReachedEventDetector(floorReachedEventChannel chan int) {
	var previousFloorSensorSignal = Elev_get_floor_sensor_signal()
	for {
		if (Elev_get_floor_sensor_signal() != -1 && previousFloorSensorSignal == -1) {
			floorReachedEventChannel <- Elev_get_floor_sensor_signal()
		}
		previousFloorSensorSignal = Elev_get_floor_sensor_signal()
		time.Sleep(10*time.Millisecond)
	}
}

func NewOrderInEmptyQueueEventDetector() bool {
	return ExactlyOneOrder()
}

func NewOrderInCurrentFloorEventDetector(order Order) bool {
	return order.Floor == Msg.PrevFloor
}

func UpDateOrders(otherLift MSG){
	switch otherLift.MsgType{
	case ADD_ORDERS:
		for i:=0;i<N_FLOORS;i++ {
			if (otherLift.ExUpOrders[i] == 1) {
				Msg.ExUpOrders[i] = 1
				Elev_set_button_lamp(BUTTON_CALL_UP, i, ON)
			}
			if( otherLift.ExDownOrders[i] == 1) {
				Msg.ExDownOrders[i] = 1
				Elev_set_button_lamp(BUTTON_CALL_DOWN, i, ON)
			}
		}
			
	case REMOVE_ORDERS:
		for i:=0;i<N_FLOORS;i++ {
			if (otherLift.ExUpOrders[i] == 0) {
				Msg.ExUpOrders[i] = 0
				Elev_set_button_lamp(BUTTON_CALL_UP, i, OFF)
			}
			if( otherLift.ExDownOrders[i] == 0) {
				Msg.ExDownOrders[i] = 0
				Elev_set_button_lamp(BUTTON_CALL_DOWN, i, OFF)
			}
		}
	}
}


func EventHandler(timerChan chan string, timeOutChan chan int, send_ch, receive_ch chan Udp_message) {
	orderEventChannel := make(chan Order)
	floorReachedEventChannel := make(chan int)
	go buttonEventDetector(orderEventChannel)
	go floorReachedEventDetector(floorReachedEventChannel)
	
	// go func(){
	// 	for {
	// 		encMsg := EncodeMsg(Msg)
	// 		Udp_msg.Data = encMsg
	// 		send_ch <- Udp_msg
	// 		fmt.Println("beat")
	// 		time.Sleep(1000*time.Millisecond)

	// 	}
		
	// }()

	for {
		
		select {

		case UDP_Rec := <- receive_ch:

			fmt.Println("HEIHEIHEHEHI", Laddr.String())

			if (Laddr.String() != UDP_Rec.Raddr) {
				fmt.Println("beat2")
				fmt.Println(UDP_Rec.Raddr)
				Dec_Msg := DecodeMsg(UDP_Rec.Data, UDP_Rec.Length)

				UpDateOrders(Dec_Msg)
				fmt.Println(Dec_Msg)
			}

		case order := <- orderEventChannel:
			AddOrder(order)
			PrintMsg()



			
			Udp_msg.Data = EncodeMsg(Msg)
			send_ch <- Udp_msg

			if (NewOrderInEmptyQueueEventDetector()) {
				fmt.Println("NewOrderInEmptyQueue")
				NewOrderInEmptyQueue(timerChan)
				fmt.Println("Event : NewOrderInEmptyQueue")
			}
			if (NewOrderInCurrentFloorEventDetector(order)) {
				NewOrderInCurrentFloor(timerChan)
				fmt.Println("Event : NewOrderInCurrentFloor")
			}

		case floor := <- floorReachedEventChannel:
			Msg.PrevFloor = floor
			fmt.Println("Event : New floor reached :", floor)
			stopped := false
			FloorReached(timerChan, stopped)
			if stopped {
				Msg.MsgType = REMOVE_ORDERS
				Udp_msg.Data = EncodeMsg(Msg)
				send_ch <- Udp_msg
				Msg.MsgType = NOTHING
			}

		case <- timerChan: 
			TimerOut()
		}
	}
}





