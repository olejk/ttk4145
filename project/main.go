package main

import (
	"fmt"
	"driver"
)

func main() {
    // Initialize hardware
    if (driver.Elev_init() == 0) {
        fmt.Printf("Unable to initialize elevator hardware!\n")
    }

    fmt.Printf("Press STOP button to stop elevator and exit program.\n");

    driver.Elev_set_motor_direction(driver.DIR_UP)

    for {
        // Change direction when we reach top/bottom floor
        if (driver.Elev_get_floor_sensor_signal() == driver.N_FLOORS - 1) {
            driver.Elev_set_motor_direction(driver.DIR_DOWN)
        } else if (driver.Elev_get_floor_sensor_signal() == 0) {
            driver.Elev_set_motor_direction(driver.DIR_UP)
        }

        // Stop elevator and exit program if the stop button is pressed
        if (driver.Elev_get_stop_signal() == 1) {
            driver.Elev_set_motor_direction(driver.DIR_STOP)
            break
        }
    }
}

