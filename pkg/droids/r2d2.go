package droids

import (
	"sync"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	gobotR2D2 "gobot.io/x/gobot/platforms/sphero/r2d2"
)

type r2d2 struct {
	table map[string]*gobotR2D2.Driver
}

func NewR2D2(table map[string]*gobotR2D2.Driver, mrb *mruby.Mrb) *r2d2 {
	droids := r2d2{
		table: table,
	}

	class := mrb.DefineClass("R2D2", nil)
	_, err := mrb.LoadString(`
class R2D2
  attr_accessor :name
end
	`)
	if err != nil {
		panic(err.Error())
	}

	class.DefineMethod("initialize", droids.Initialize, mruby.ArgsReq(1))
	class.DefineMethod("dome", droids.Dome, mruby.ArgsReq(1))
	class.DefineMethod("tripod", droids.Tripod, mruby.ArgsReq(0))
	class.DefineMethod("bipod", droids.Bipod, mruby.ArgsReq(0))
	class.DefineMethod("macro", droids.Macro, mruby.ArgsReq(1))
	class.DefineMethod("move", droids.Move, mruby.ArgsReq(2))

	return &droids
}

func (droids *r2d2) driver(mrbInstance *mruby.MrbValue) *gobotR2D2.Driver {
	value, err := mrbInstance.Call("name")
	var driver *gobotR2D2.Driver

	if err != nil {
		panic(err.Error())
	} else {
		ident := string(value.String())
		driver = droids.table[ident]
	}

	return driver
}

func (droids *r2d2) Initialize(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	name := string(args[0].String())
	bleAdaptor := ble.NewClientAdaptor(name)
	droidDriver := gobotR2D2.NewDriver(bleAdaptor)
	droids.table[name] = droidDriver
	self.Call("name=", args[0])

	var wg sync.WaitGroup
	wg.Add(1)
	work := func() {
		wg.Done()
	}

	robot := gobot.NewRobot("R2D2",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{droidDriver},
		work,
	)

	go robot.Start()
	wg.Wait()

	return self, nil
}

func (droids *r2d2) Dome(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	heading := int16(args[0].Fixnum())

	droids.driver(self).Dome(heading)

	return nil, nil
}

func (droids *r2d2) Tripod(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	droids.driver(self).Tripod()

	return nil, nil
}

func (droids *r2d2) Bipod(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	droids.driver(self).Bipod()

	return nil, nil
}

func (droids *r2d2) Macro(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	num := uint8(args[0].Fixnum())
	droids.driver(self).Macro(num)

	return nil, nil
}

func (droids *r2d2) Move(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	speed := uint8(args[0].Fixnum())
	heading := uint16(args[1].Fixnum())
	droids.driver(self).Move(speed, heading)

	return nil, nil
}
