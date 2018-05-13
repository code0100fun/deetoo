package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot/platforms/sphero/r2q5"

	"github.com/hone/mrgoboto/droids"
)

func startWorker(file string, droidTable map[string]*r2q5.Driver, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		mrb := mruby.NewMrb()
		defer mrb.Close()
		defer wg.Done()

		droids.NewR2D2(droidTable, mrb)

		dat, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		content := string(dat)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			mrb.Close()
			wg.Done()
		}()
		result, err := mrb.LoadString(content)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Done: %s\n", result)
	}()
}

func main() {
	r2d2Table := make(map[string]*r2q5.Driver)

	var wg sync.WaitGroup
	startWorker("main.mrb", r2d2Table, &wg)
	wg.Wait()
}
