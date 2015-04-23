package queue

import (
	. "def"
)

func ShouldIStop() {
	
}

func NextDirection(msg Msg) {
	if(msg.InOrders[msg.Dir] == 1){
		return DIR_STOP
	}

	if(msg.Dir == UP){
		for i:=msg.Dir+1;i<N_FLOORS;i++{
			if(msg.InOrders[msg.Dir] == 1){
				return msg.Dir
			}
		}

		

	}
	

	if(Msg.Dir==UP){

	}

	
}