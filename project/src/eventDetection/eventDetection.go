package eventDetection

import (
	"fmt"
	"time"
	. "def"
	. "driver"
	. "stateMachine"
	. "queue"
)

func buttonEventDetector(orderEventChannel chan Order){
	var currentSignalMatrix 	[3][N_FLOORS]int 
	var previousSignalMatrix 	[3][N_FLOORS]int 
	
	for{
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
	for{
		if (Elev_get_floor_sensor_signal() != -1 && previousFloorSensorSignal == -1) {
			fmt.Println("7")
			floorReachedEventChannel <- Elev_get_floor_sensor_signal()
			fmt.Println("77")
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


func EventHandler(timerChan chan string, timeOutChan chan int) {
	orderEventChannel := make(chan Order)
	floorReachedEventChannel := make(chan int)
	go buttonEventDetector(orderEventChannel)
	go floorReachedEventDetector(floorReachedEventChannel)
	
	for{
		select{
		case order := <- orderEventChannel:
			AddOrder(order)
			PrintMsg()
			if(NewOrderInEmptyQueueEventDetector()) {
				fmt.Println("NewOrderInEmptyQueue")
				NewOrderInEmptyQueue(timerChan)
				fmt.Println("Event : NewOrderInEmptyQueue")
			}
			if(NewOrderInCurrentFloorEventDetector(order)) {
				NewOrderInCurrentFloor(timerChan)
				fmt.Println("Event : NewOrderInCurrentFloor")
			}

		case floor := <- floorReachedEventChannel:
			Msg.PrevFloor = floor
			fmt.Println("Event : New floor reached :", floor)
			FloorReached(timerChan)

		case <- timerChan: 
			TimerOut()
		}
	}
}





