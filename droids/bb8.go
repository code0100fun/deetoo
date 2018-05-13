package droids

import (
	"sync"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	gobotBB8 "gobot.io/x/gobot/platforms/sphero/bb8"
)

type bb8 struct {
	table map[string]*gobotBB8.BB8Driver
}

func NewBB8(table map[string]*gobotBB8.BB8Driver, mrb *mruby.Mrb) *bb8 {
	droids := bb8{
		table: table,
	}

	class := mrb.DefineClass("BB8", nil)
	_, err := mrb.LoadString(`
class BB8
  attr_accessor :name
end
	`)
	if err != nil {
		panic(err.Error())
	}
	class.DefineMethod("initialize", droids.Initialize, mruby.ArgsReq(1))
	class.DefineMethod("set_rgb", droids.SetRGB, mruby.ArgsReq(3))
	class.DefineMethod("roll", droids.SetRGB, mruby.ArgsReq(2))
	class.DefineMethod("boost", droids.Boost, mruby.ArgsReq(1))
	class.DefineMethod("set_rotation_rate", droids.SetRotationRate, mruby.ArgsReq(1))
	class.DefineMethod("set_stabilization", droids.SetStabilization, mruby.ArgsReq(1))
	class.DefineMethod("set_back_led_output", droids.SetBackLEDOutput, mruby.ArgsReq(1))

	return &droids
}

func (droidss *bb8) driver(mrbInstance *mruby.MrbValue) *gobotBB8.BB8Driver {
	value, err := mrbInstance.Call("name")
	var driver *gobotBB8.BB8Driver

	if err != nil {
		panic(err.Error())
	} else {
		ident := string(value.String())
		driver = droidss.table[ident]
	}

	return driver
}

func (droidss *bb8) Initialize(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	name := string(args[0].String())
	bleAdaptor := ble.NewClientAdaptor(name)
	droidsDriver := gobotBB8.NewDriver(bleAdaptor)
	droidss.table[name] = droidsDriver
	self.Call("name=", args[0])

	var wg sync.WaitGroup
	wg.Add(1)
	work := func() {
		wg.Done()
	}

	robot := gobot.NewRobot("BB8",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{droidsDriver},
		work,
	)

	go robot.Start()
	wg.Wait()

	return self, nil
}

func (droids *bb8) SetRGB(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	r := uint8(args[0].Fixnum())
	g := uint8(args[1].Fixnum())
	b := uint8(args[2].Fixnum())

	droids.driver(self).SetRGB(r, g, b)

	return nil, nil
}

func (droids *bb8) Roll(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	speed := uint8(args[0].Fixnum())
	heading := uint16(args[1].Fixnum())

	droids.driver(self).Roll(speed, heading)

	return nil, nil
}

func (droids *bb8) Boost(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	var state bool

	if args[0].Type() == mruby.TypeFalse || args[0].Type() == mruby.TypeNil {
		state = false
	} else {
		state = true
	}

	droids.driver(self).Boost(state)

	return nil, nil
}

func (droids *bb8) SetRotationRate(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	speed := uint8(args[0].Fixnum())

	droids.driver(self).SetRotationRate(speed)

	return nil, nil
}

func (droids *bb8) SetStabilization(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	var state bool

	if args[0].Type() == mruby.TypeFalse || args[0].Type() == mruby.TypeNil {
		state = false
	} else {
		state = true
	}

	droids.driver(self).SetStabilization(state)

	return nil, nil
}

func (droids *bb8) SetBackLEDOutput(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	value := uint8(args[0].Fixnum())

	droids.driver(self).SetBackLEDOutput(value)

	return nil, nil
}
