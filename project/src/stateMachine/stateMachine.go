package stateMachine

import (
	. "driver"
	"fmt"
	"queue"
	. "def"
) 

func floorReached(timerChan chan string){
	switch State{
		case IDLE: 
			fmt.Println("Error: no floor reached in state IDLE")
		case MOVING:
			if(ShouldIStop){										//må skrive hvilke pakke funksjon er i. Funksjon ikke implementert
				Elev_set_motor_direction(DIR_STOP)
				Elev_set_door_open_lamp(ON)
				timerChan <- "START"
				State = DOOR_OPEN
			}
		case DOOR_OPEN:
			fmt.Println("Error: no floor reached in state DOOR_OPEN")
	}
}

func timerOut(){
	switch State{
		case IDLE:
			fmt.Println("ERROR: Timeout in state IDLE")
		case MOVING:
			fmt.Println("ERROR: Timeout in state MOVING")
		case DOOR_OPEN:
			Elev_set_door_open_lamp(OFF)
			if(nextDirection() == DIR_STOP ){
				State = IDLE
				return 
			}else if( nextDirection() != message.Dir){
				message.Dir = nextDirection()
			}
			Elev_set_motor_direction(message.Dir)
	}
}

func newOrderInEmptyQueue(doorTimer chan int){
	switch State{
	case IDLE:
		if(nextDirection() == DIR_STOP){
			Elev_set_door_open_lamp(ON)
			doorTimer <- START
			State = DOOR_OPEN
		} else {
			Elev_set_motor_direction(nextDirection())
			State = MOVING
		}
	case MOVING:
		fmt.Println("Error: newOrderInEmptyQueue in state MOVING")
	case DOOR_OPEN:
		fmt.Println("Error: newOrderInEmptyQueue in state DOOR_OPEN")
	}
}

func newOrderInCurrentFloor(timerChan chan int) {
	case IDLE:
		Elev_set_door_open_lamp(ON)
		State = DOOR_OPEN
		timerChan <- "START"
	case MOVING:
		fmt.Println("Error: New order in current floor while moving.")
	case DOOR_OPEN:
		timerChan <- "START"
}



