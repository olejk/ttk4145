package queue

import (
	"fmt"
	. "def"
	//"time"
	. "driver"
)

func ZeroOrders() bool {
	for i:=0;i<N_FLOORS;i++ {
		//fmt.Println(Msg.ExUpOrders[i], " ", Msg.ExDownOrders[i], " ", Msg.InOrders[i])
		if (Msg.ExUpOrders[i] == 1 || Msg.ExDownOrders[i] == 1 || Msg.InOrders[i] == 1 ) {
			return false
		}
	}
	return true
}

func ExactlyOneOrder() bool {
	sum :=0
	for i:=0;i<N_FLOORS;i++ {
		//fmt.Println(Msg.ExUpOrders[i], " ", Msg.ExDownOrders[i], " ", Msg.InOrders[i])
		if (Msg.ExUpOrders[i] == 1) {
			sum++
		} 
		if (Msg.ExDownOrders[i] == 1 ) {
			sum++
		}
		if (Msg.InOrders[i] == 1 ) {
			sum++
		}
	}
	return sum == 1
}

// 1:  first checks if it should stop in the current floor
// 2:  if not, check orders furter in the current direction 
// 3:  if not, check orders in oppisite direction in current floor
// 4:  if not, change direction 
func NextDirection() int {
	fmt.Println("Msg.Dir = ", Msg.Dir)
	if (Msg.InOrders[Msg.PrevFloor] == 1) {
		fmt.Println("NextDirection(): a")
		return DIR_STOP
	}
	if (Msg.Dir == DIR_UP && Msg.ExUpOrders[Msg.PrevFloor] == 1 || Msg.Dir == DIR_DOWN && 
	Msg.ExDownOrders[Msg.PrevFloor] == 1) {
		fmt.Println("NextDirection(): b")
		return DIR_STOP	
	}
	if (ZeroOrders()) {
		fmt.Println("NextDirection(): c")
		return DIR_STOP
	}
	if (Msg.Dir == DIR_UP) {
		for i:=Msg.PrevFloor+1;i<N_FLOORS;i++ {
			if(Msg.InOrders[i] == 1 || Msg.ExUpOrders[i] == 1 || Msg.ExDownOrders[i] == 1) {
				fmt.Println("NextDirection(): d")
				return DIR_UP
			}
		}
		if (Msg.ExDownOrders[Msg.PrevFloor]  == 1) {
			fmt.Println("NextDirection(): e")
			return DIR_STOP
		}
		fmt.Println("ERROR: error in NextDirection-function, with Msg.Dir UP" )
		return DIR_DOWN
					
	} else {
		for i:=Msg.PrevFloor-1;i>=0;i-- {
			if (Msg.InOrders[i] == 1 || Msg.ExUpOrders[i] == 1 || Msg.ExDownOrders[i] == 1) {
				fmt.Println("NextDirection(): f")
				return DIR_DOWN
			}
		}
		if (Msg.ExUpOrders[Msg.PrevFloor] == 1) {
			fmt.Println("NextDirection(): g")
			return DIR_STOP
		}
		fmt.Println("ERROR: error in NextDirection-function, with Msg.Dir DOWN")
		return DIR_UP
	}		
}

func ExistOrdersInCurrentDir() bool {
	if (Msg.InOrders[Msg.PrevFloor] == 1) {
			return true
	}
	if (Msg.Dir == DIR_UP) {
		for i:=Msg.PrevFloor+1;i<N_FLOORS;i++ {
			if(Msg.InOrders[i] == 1 || Msg.ExUpOrders[i] == 1 || Msg.ExDownOrders[i] == 1) {
				fmt.Println("ExistOrdersInCurrentDir: 1")
				return true
			}
		}

		if(Msg.ExUpOrders[Msg.PrevFloor] == 1){
			return true
		}
	} else if (Msg.Dir == DIR_DOWN) {
		for i:=Msg.PrevFloor-1;i>=0;i-- {
			if (Msg.InOrders[i] == 1 || Msg.ExUpOrders[i] == 1 || Msg.ExDownOrders[i] == 1) {
				fmt.Println("ExistOrdersInCurrentDir: 2")
				return true
			}
		}
		if(Msg.ExDownOrders[Msg.PrevFloor] == 1){
			return true
		}
	}
	return false
}


func RemoveDoneOrders() {
	Msg.InOrders[Msg.PrevFloor] = 0
	Elev_set_button_lamp(BUTTON_COMMAND, Msg.PrevFloor, OFF)
	if (Msg.Dir == DIR_UP) {
		Msg.ExUpOrders[Msg.PrevFloor] = 0 
		Elev_set_button_lamp(BUTTON_CALL_UP, Msg.PrevFloor, OFF)	
	} else if (Msg.Dir == DIR_DOWN) {
		Msg.ExDownOrders[Msg.PrevFloor] = 0
		Elev_set_button_lamp(BUTTON_CALL_DOWN, Msg.PrevFloor, OFF)
	}
	if (NextDirection() != Msg.Dir) {
		fmt.Println("ND ", NextDirection())
		fmt.Println("Msg.Dir: ", Msg.Dir)
		fmt.Println("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
		if (Msg.Dir == DIR_UP) {
			Msg.ExDownOrders[Msg.PrevFloor] = 0 
			Elev_set_button_lamp(BUTTON_CALL_DOWN, Msg.PrevFloor, OFF)	
		} else if (Msg.Dir == DIR_DOWN) {
			Msg.ExUpOrders[Msg.PrevFloor] = 0
			Elev_set_button_lamp(BUTTON_CALL_UP, Msg.PrevFloor, OFF)
		}
	}
}

//Funksjonen burde kanskje flyttes til en annen package
func AddOrder(order Order) {
	switch order.Button {
	case BUTTON_CALL_UP: 
		Msg.ExUpOrders[order.Floor] = 1
	case BUTTON_CALL_DOWN:
		Msg.ExDownOrders[order.Floor] = 1
	case BUTTON_COMMAND:
		Msg.InOrders[order.Floor] = 1
	}
	//Bør kanskje flyttes
	Elev_set_button_lamp(order.Button, order.Floor, ON)	

	Msg.MsgType = ADD_ORDERS
}







