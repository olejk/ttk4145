package timer

import(
	"fmt"
	"time"
)

const(
	DOOR_OPEN_TIME = 3*time.Second
)

var(
	timer 		time.Time
	timerFlag 	bool 		= false
)

func checkTimer(timeOutChan chan int) {
	for {
		if (time.Since(timer) > DOOR_OPEN_TIME) && timerFlag == true {
			timerFlag = false
			timeOutChan <- 0
			fmt.Println("Timeout")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func DoorTimer(timerChan chan string, timeOutChan chan int) {
	go checkTimer(timeOutChan)

	for {
		select {
			case <- timerChan:
				fmt.Println("Starting timer")
				timer = time.Now()
				timerFlag = true
			case <- timeOutChan:
				timerChan <- "timeout"
		}
	}
}

