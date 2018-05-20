package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot/platforms/sphero/bb8"
	"gobot.io/x/gobot/platforms/sphero/r2q5"

	"github.com/code0100fun/mrgoboto/droids"
)

type droidTable struct {
	bb8  map[string]*bb8.BB8Driver
	r2d2 map[string]*r2q5.Driver
}

func NewDroidTable() droidTable {
	table := droidTable{
		bb8:  make(map[string]*bb8.BB8Driver),
		r2d2: make(map[string]*r2q5.Driver),
	}

	return table
}

func startWorker(file string, table droidTable, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		mrb := mruby.NewMrb()
		defer mrb.Close()

		droids.NewR2D2(table.r2d2, mrb)
		droids.NewBB8(table.bb8, mrb)

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
	var wg sync.WaitGroup
	table := NewDroidTable()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		wg.Done()
		wg.Done()
	}()

	startWorker("main.mrb", table, &wg)
	startWorker("bb8.mrb", table, &wg)
	wg.Wait()
}
