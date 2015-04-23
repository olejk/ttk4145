package def

const (

	N_ELEV				int = 1


	N_FLOORS			int = 4
	N_BUTTONS			int = 3
	BUTTON_CALL_UP		int = 0
	BUTTON_CALL_DOWN	int = 1
	BUTTON_COMMAND		int = 2
	DIR_UP				int = 1
	DIR_DOWN			int = -1
	DIR_STOP			int = 0
	ON 					int = 1
	OFF					int = 0

	//states
	IDLE 		int = 0
	MOVING 		int = 1
	DOOR_OPEN 	int = 2

)

type Msg struct{
	State 			int
	PrevFloor 		int
	Dir   			int 	//never 0. 
	ExUpOrders 		[N_FLOORS]int
	ExDownOrders	[N_FLOORS]int
	InOrders		[N_FLOORS]int

}
