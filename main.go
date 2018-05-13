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

	"github.com/hone/mrgoboto/droids"
)

func startWorker(file string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		mrb := mruby.NewMrb()
		defer mrb.Close()

		r2d2Table := make(map[string]*r2q5.Driver)
		droids.NewR2D2(r2d2Table, mrb)

		bb8Table := make(map[string]*bb8.BB8Driver)
		droids.NewBB8(bb8Table, mrb)

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		wg.Done()
		wg.Done()
	}()

	startWorker("main.mrb", &wg)
	startWorker("bb8.mrb", &wg)
	wg.Wait()
}
