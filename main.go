package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "sync"
    //"time"

    "github.com/mitchellh/go-mruby"

    "gobot.io/x/gobot"
    "gobot.io/x/gobot/platforms/ble"
    "gobot.io/x/gobot/platforms/sphero/r2q5"
)

func startWorker(r2q5 *r2q5.Driver, file string, wg *sync.WaitGroup) {
    wg.Add(1)
    go func() {
        defer wg.Done()
        mrb := mruby.NewMrb()
        defer mrb.Close()

	NewMrubyR2q5(r2q5, mrb)

        dat, err := ioutil.ReadFile(file)
        if err != nil {
            panic(err)
        }
        content := string(dat)

        result, err := mrb.LoadString(content)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Done: %s\n", result)
    }()
}

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	r2q5 := r2q5.NewDriver(bleAdaptor)
	defer r2q5.Sleep()

	work := func() {
		var wg sync.WaitGroup
		startWorker(r2q5, "main.mrb", &wg)
		wg.Wait()
	}

	robot := gobot.NewRobot("R2Q5",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{r2q5},
		work,
	)

	robot.Start()
}
