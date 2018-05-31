// +build example
//
// Do not build by default.

/*
 How to run
 Pass the Bluetooth address or name as the first param:

	go run examples/r2d2.go R2-1234

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/r2d2"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	r2d2 := r2d2.NewDriver(bleAdaptor)

	work := func() {
		gobot.Every(4000*time.Millisecond, func() {
			r2d2.Tripod()
			time.Sleep(2000*time.Millisecond)
			r2d2.Bipod()
		})
	}

	robot := gobot.NewRobot("R2D2",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{r2d2},
		work,
	)

	robot.Start()
}
