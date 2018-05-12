package droids

import (
	"github.com/mitchellh/go-mruby"

	gobotBB8 "gobot.io/x/gobot/platforms/sphero/bb8"
)

type bb8 struct {
	driver *gobotBB8.BB8Driver
	mrb    *mruby.Mrb
}

func NewBB8(driver *gobotBB8.BB8Driver, mrb *mruby.Mrb) *bb8 {
	droid := bb8{
		driver: driver,
		mrb:    mrb,
	}

	class := mrb.DefineClass("BB8", nil)
	class.DefineMethod("set_rgb", droid.SetRGB, mruby.ArgsReq(3))
	class.DefineMethod("roll", droid.SetRGB, mruby.ArgsReq(2))
	class.DefineMethod("boost", droid.Boost, mruby.ArgsReq(1))
	class.DefineMethod("set_rotation_rate", droid.SetRotationRate, mruby.ArgsReq(1))
	class.DefineMethod("set_stabilization", droid.SetStabilization, mruby.ArgsReq(1))
	class.DefineMethod("set_back_led_output", droid.SetBackLEDOutput, mruby.ArgsReq(1))

	return &droid
}

func (droid *bb8) SetRGB(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	r := uint8(args[0].Fixnum())
	g := uint8(args[1].Fixnum())
	b := uint8(args[2].Fixnum())

	droid.driver.SetRGB(r, g, b)

	return nil, nil
}

func (droid *bb8) Roll(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	speed := uint8(args[0].Fixnum())
	heading := uint16(args[1].Fixnum())

	droid.driver.Roll(speed, heading)

	return nil, nil
}

func (droid *bb8) Boost(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	var state bool

	if args[0].Type() == mruby.TypeFalse || args[0].Type() == mruby.TypeNil {
		state = false
	} else {
		state = true
	}

	droid.driver.Boost(state)

	return nil, nil
}

func (droid *bb8) SetRotationRate(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	speed := uint8(args[0].Fixnum())

	droid.driver.SetRotationRate(speed)

	return nil, nil
}

func (droid *bb8) SetStabilization(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	var state bool

	if args[0].Type() == mruby.TypeFalse || args[0].Type() == mruby.TypeNil {
		state = false
	} else {
		state = true
	}

	droid.driver.SetStabilization(state)

	return nil, nil
}

func (droid *bb8) SetBackLEDOutput(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	value := uint8(args[0].Fixnum())

	droid.driver.SetBackLEDOutput(value)

	return nil, nil
}
