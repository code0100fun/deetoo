package droids

import (
	"github.com/mitchellh/go-mruby"

	gobotR2Q5 "gobot.io/x/gobot/platforms/sphero/r2q5"
)

type r2d2 struct {
	driver *gobotR2Q5.Driver
	mrb    *mruby.Mrb
}

func NewR2D2(driver *gobotR2Q5.Driver, mrb *mruby.Mrb) *r2d2 {
	droid := r2d2{
		driver: driver,
		mrb:    mrb,
	}

	class := mrb.DefineClass("R2D2", nil)
	class.DefineMethod("dome", droid.Dome, mruby.ArgsReq(1))
	class.DefineMethod("tripod", droid.Tripod, mruby.ArgsReq(0))
	class.DefineMethod("bipod", droid.Bipod, mruby.ArgsReq(0))
	class.DefineMethod("macro", droid.Macro, mruby.ArgsReq(1))
	class.DefineMethod("move", droid.Move, mruby.ArgsReq(2))

	return &droid
}

func (droid *r2d2) Dome(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	heading := int16(args[0].Fixnum())
	droid.driver.Dome(heading)

	return nil, nil
}

func (droid *r2d2) Tripod(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	droid.driver.Tripod()

	return nil, nil
}

func (droid *r2d2) Bipod(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	droid.driver.Bipod()

	return nil, nil
}

func (droid *r2d2) Macro(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	num := uint8(args[0].Fixnum())
	droid.driver.Macro(num)

	return nil, nil
}

func (droid *r2d2) Move(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := droid.mrb.GetArgs()
	speed := uint8(args[0].Fixnum())
	heading := uint16(args[1].Fixnum())
	droid.driver.Move(speed, heading)

	return nil, nil
}
