package driver

import()

const (
	N_FLOORS			int = 4
	N_BUTTONS			int = 3
	BUTTON_CALL_UP		int = 0
	BUTTON_CALL_DOWN	int = 1
	BUTTON_COMMAND		int = 2
	DIR_UP				int = 1
	DIR_DOWN			int = -1
	DIR_STOP			int = 0
)

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

func Elev_init() int {
	// Init hardware
	if Io_init() == 0 {
		return 0
	}
	// Zero all floor button lamps
	for i := 0; i < N_FLOORS; i++ {
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

func Elev_set_motor_direction(dir int) {
	// Legge til fart?
	if dir == 0 {
		Io_write_analog(MOTOR, 0)
	} else if dir > 0 {
		Io_clear_bit(MOTORDIR)
		Io_write_analog(MOTOR, 2800)
	} else if dir < 0 {
		Io_set_bit(MOTORDIR)
		Io_write_analog(MOTOR, 2800)
	}
}

func Elev_get_obstruction_signal() int {
	return Io_read_bit(OBSTRUCTION)
}

func Elev_get_stop_signal() int {
	return Io_read_bit(STOP)
}

func Elev_set_stop_lamp(value int) {
	if value != 0 {
		Io_set_bit(LIGHT_STOP)
	}
	if value == 0 {
		Io_clear_bit(LIGHT_STOP)
	}
}

func Elev_set_door_open_lamp(value int) {
	if value == 0 {
		Io_clear_bit(LIGHT_DOOR_OPEN)
	} else {
		Io_set_bit(LIGHT_DOOR_OPEN)
	} 
}


func Elev_get_floor_sensor_signal() int {
	if Io_read_bit(SENSOR_FLOOR1) != 0 {
		return 0
	} else if Io_read_bit(SENSOR_FLOOR2) != 0 {
        return 1
    } else if Io_read_bit(SENSOR_FLOOR3) != 0 {
        return 2
    } else if Io_read_bit(SENSOR_FLOOR4) != 0 {
        return 3
    } else {
        return -1
    }
}


func Elev_set_floor_indicator(floor int) int {
	if floor < 0 || floor >= N_FLOORS {
		return -1
		//abort
	}

	if (floor & 0x02) != 0 {
		Io_set_bit(LIGHT_FLOOR_IND1)
	} else {
		Io_clear_bit(LIGHT_FLOOR_IND1)
	}
	if (floor & 0x01) != 0 {
		Io_set_bit(LIGHT_FLOOR_IND2)
	} else {
		Io_clear_bit(LIGHT_FLOOR_IND2)
	}
	return 1
}

func Elev_get_button_signal(button int, floor int) int {
	if floor < 0 || floor >= N_FLOORS {
		return -1
		//blablalblabl
	}
	if (button == BUTTON_CALL_UP && floor == N_FLOORS - 1) || (button == BUTTON_CALL_DOWN && floor == 0) {
		//ingen oppknapp i øverste etasje, og heller ingen nedknapp i nederste etasje
	}
	if !(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND) {
		//ugyldig verdi
	}

	if Io_read_bit(button_channel_matrix[floor][button]) != 0 {
		return 1
	} else {
		return 0
	}
}

func Elev_set_button_lamp(button int, floor int, value int) int {
	if floor < 0 || floor >= N_FLOORS {
		return -1
		//blablalblabl
	}
	if (button == BUTTON_CALL_UP && floor == N_FLOORS - 1) || (button == BUTTON_CALL_DOWN && floor == 0) {
		//ingen oppknapp i øverste etasje, og heller ingen nedknapp i nederste etasje
	}
	if !(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND) {
		//ugyldig verdi
	}

	if value != 0 {
		Io_set_bit(lamp_channel_matrix[floor][button])
	} else {
		Io_clear_bit(lamp_channel_matrix[floor][button])
	}

	return 1
}
