package stateMachine

import (
	. "driver"
	"fmt"
	. "def"
	. "queue"
)

func FloorReached(timerChan chan string, stopped bool) {
	Elev_set_floor_indicator(Elev_get_floor_sensor_signal())
	switch Msg.State {
		case IDLE: 
			fmt.Println("Error: no floor reached in state IDLE")
		case MOVING:
			if (NextDirection() == 0) {
				stopped = true
				fmt.Println("FloorReached: 1")
				RemoveDoneOrders()						
				Elev_set_motor_direction(DIR_STOP)
				Elev_set_door_open_lamp(ON)
				timerChan <- "START"
				Msg.State = DOOR_OPEN
			}
		case DOOR_OPEN:
			fmt.Println("Error: no floor reached in state DOOR_OPEN")
	}
}

func TimerOut() {
	switch Msg.State {
		case IDLE:
			fmt.Println("ERROR: Timeout in state IDLE")
		case MOVING:
			fmt.Println("ERROR: Timeout in state MOVING")
		case DOOR_OPEN:
			Elev_set_door_open_lamp(OFF)
			if (NextDirection() == DIR_STOP ) {
				fmt.Println("TimerOut: 1")
				Msg.State = IDLE
				return 
			} else if (NextDirection() != Msg.Dir) {
				fmt.Println("TimerOut: 2")
				Msg.Dir = NextDirection()
				RemoveDoneOrders()
			}
			Elev_set_motor_direction(Msg.Dir)
			Msg.State = MOVING
	}
}

func NewOrderInEmptyQueue(timerChan chan string) {
	fmt.Println("NewOrderInEmptyQueue: 1")
	fmt.Println(NextDirection())
	switch Msg.State {
	case IDLE:
		if (NextDirection() == DIR_STOP) {
			Elev_set_door_open_lamp(ON)
			RemoveDoneOrders()
			fmt.Println("NewOrderInEmptyQueue: 2")
			timerChan <- "START"	
			fmt.Println("NewOrderInEmptyQueue: 3")
			Msg.State = DOOR_OPEN
		} else {
			Elev_set_motor_direction(NextDirection())
			fmt.Println("NewOrderInEmptyQueue: 4")
			// Msg.Dir = NextDirection() 
			Msg.State = MOVING
		}
	case MOVING:
		fmt.Println("Error: newOrderInEmptyQueue in state MOVING")
	case DOOR_OPEN:
		fmt.Println("Error: newOrderInEmptyQueue in state DOOR_OPEN")
	}
}

func NewOrderInCurrentFloor(timerChan chan string) {
	switch Msg.State{
		case IDLE:
			RemoveDoneOrders()
			Elev_set_door_open_lamp(ON)
			Msg.State = DOOR_OPEN
			timerChan <- "START"
		case MOVING:
			fmt.Println("Error: New order in current floor while moving.")
		case DOOR_OPEN:
			RemoveDoneOrders()
			timerChan <- "START"
	}
}



