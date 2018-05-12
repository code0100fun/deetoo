package main

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot/platforms/sphero/r2q5"

	"github.com/hone/mrgoboto/droids"
)

func startWorker(file string, droidTable map[string]*r2q5.Driver, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		mrb := mruby.NewMrb()
		defer wg.Done()

		droids.NewR2D2(droidTable, mrb)

		dat, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		content := string(dat)

		result, err := mrb.LoadString(content)
		if err != nil {
			panic(err)
		}
		mrb.Close()
		fmt.Printf("Done: %s\n", result)
	}()
}

func main() {
	r2d2Table := make(map[string]*r2q5.Driver)

	var wg sync.WaitGroup
	startWorker("main.mrb", r2d2Table, &wg)
	wg.Wait()
}
