package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"sync"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/r2q5"
)

func apiSendMessage(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := m.GetArgs()
	channel := args[0].String()
	message := args[1].Hash()
	fmt.Printf("Channel: %s, Message: %s\n", channel, message)
	return nil, nil
}


func startWorker(r2q5 *r2q5.Driver, file string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		mrb := mruby.NewMrb()
		defer mrb.Close()

		class := mrb.DefineClass("Go", nil)
		class.DefineClassMethod("send_message", apiSendMessage, mruby.ArgsReq(2))
		class.DefineClassMethod("dome", func(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
			args := m.GetArgs()
			heading := int16(args[0].Fixnum())
			r2q5.Dome(heading)

			return nil, nil
		}, mruby.ArgsReq(1))

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
