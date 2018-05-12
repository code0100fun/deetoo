package main

import (
	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot/platforms/sphero/r2q5"
)

type mruby_r2q5 struct {
	driver	*r2q5.Driver
	mrb	*mruby.Mrb
}

func NewMrubyR2q5(driver *r2q5.Driver, mrb *mruby.Mrb) *mruby_r2q5 {
	droid := mruby_r2q5 {
		driver: driver,
		mrb: mrb,
	}

	class := mrb.DefineClass("R2Q5", nil)
	class.DefineMethod("dome", droid.Dome, mruby.ArgsReq(1))
	class.DefineMethod("tripod", droid.Tripod, mruby.ArgsReq(0))
	class.DefineMethod("bipod", droid.Bipod, mruby.ArgsReq(0))
	class.DefineMethod("macro", droid.Macro, mruby.ArgsReq(1))
	class.DefineMethod("move", droid.Move, mruby.ArgsReq(2))

	return &droid
}

func (droid *mruby_r2q5) Dome(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	heading := int16(args[0].Fixnum())
	droid.driver.Dome(heading)

	return nil, nil
}

func (droid *mruby_r2q5) Tripod(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	droid.driver.Tripod()

	return nil, nil
}

func (droid *mruby_r2q5) Bipod(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	droid.driver.Bipod()

	return nil, nil
}

func (droid *mruby_r2q5) Macro(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	num := uint8(args[0].Fixnum())
	droid.driver.Macro(num)

	return nil, nil
}

func (droid *mruby_r2q5) Move(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	speed := uint8(args[0].Fixnum())
	heading := uint16(args[1].Fixnum())
	droid.driver.Move(speed, heading)

	return nil, nil
}
