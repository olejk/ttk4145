package driver

import {
	
}

const {
	N_FLOORS			int = 4
	BUTTON_CALL_UP		int = 0
	BUTTON_CALL_DOWN	int = 1
	BUTTON_COMMAND		int = 2
	UP_DIR				int = 1
	DOWN_DIR			int = -1
	STOP_DIR			int = 0
}

var(
	lamp_channel_matrix [N_FLOORS][N_BUTTONS] int = [N_FLOORS][N_BUTTONS] int {
    	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
	}

	button_channel_matrix [N_FLOORS][N_BUTTONS] int = [N_FLOORS][N_BUTTONS] int {
    	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
	}
) 





func Elev_init() {
	if Io_init == 0 {
		return 0
	}
	for i := 0; i < N_FLOORS; i++ {
		
	}
}

func Elev_set_motor_direction() {
	
}

func Elev_get_obstruction_signal() {
	
}

func Elev_get_stop_signal() {
	
}

func Set_stop_lamp() {
	
}

func Elev_get_floor_sensor_signal() {
	
}

func Set_floor_indicator() {
	
}

func Elev_get_button_signal() {
	
}

typedef enum tag_elev_motor_direction { 
    DIRN_DOWN = -1,
    DIRN_STOP = 0,
    DIRN_UP = 1
} elev_motor_direction_t;

typedef enum tag_elev_lamp_type { 
    BUTTON_CALL_UP = 0,
    BUTTON_CALL_DOWN = 1,
    BUTTON_COMMAND = 2
} elev_button_type_t;
