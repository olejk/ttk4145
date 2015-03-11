package driver

import {
	
}

const {
	N_FLOORS			int = 4
	N_BUTTONS			int = 3
	BUTTON_CALL_UP		int = 0
	BUTTON_CALL_DOWN	int = 1
	BUTTON_COMMAND		int = 2
	DIR_UP				int = 1
	DIR_DOWN			int = -1
	DIR_STOP			int = 0
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
	// Init hardware
	if Io_init == 0 {
		return 0
	}
	// Zero all floor button lamps
	for i := 0; i < N_FLOORS; ++i {
		if i != 0 {
			Elev_set_button_lamp(BUTTON_CALL_DOWN, i, 0)
		}
		if i != N_FLOORS-1 {
			Elev_set_button_lamp(BUTTON_CALL_UP, i, 0)
		}
		Elev_set_button_lamp(BUTTON_COMMAND, i, 0)
	}
	// Clear stop lamp, door open lamp, and set floor indicator to ground floor
	Elev_set_stop_lamp(0)
    Elev_set_door_open_lamp(0)
    Elev_set_floor_indicator(0)

    return 1
}

func Elev_set_motor_direction() {
	
}

func Elev_get_obstruction_signal() {
	
}

func Elev_get_stop_signal() {
	
}

func Elev_set_stop_lamp() {
	
}

func Elev_set_door_open_lamp() {
	
}

func Elev_get_floor_sensor_signal() {
	
}

func Elev_set_floor_indicator() {
	
}

func Elev_get_button_signal() {
	
}

func Elev_set_button_lamp(button int, floor int, value int) {
	if value != 0 {
		Io_set_bit(lamp_channel_matrix[floor][button])
	} else {
		Io_clear_bit(lamp_channel_matrix[floor][button])
	}
}
