package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mitchellh/go-mruby"

	"gobot.io/x/gobot/platforms/sphero/bb8"
	"gobot.io/x/gobot/platforms/sphero/r2d2"

	"github.com/code0100fun/deetoo/pkg/droids"
)

type droidTable struct {
	bb8  map[string]*bb8.BB8Driver
	r2d2 map[string]*r2d2.Driver
}

func NewDroidTable() droidTable {
	table := droidTable{
		bb8:  make(map[string]*bb8.BB8Driver),
		r2d2: make(map[string]*r2d2.Driver),
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

		if _, err := os.Stat(file); os.IsNotExist(err) {
			os.Stderr.WriteString(fmt.Sprintf("%s does not exist.\n", file))
			wg.Done()
			return
		}

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
		cleanup(table, wg)
	}()
}

func sleepAllDroids(table droidTable) {
	for _, droid := range table.r2d2 {
		droid.Sleep()
	}
	for _, droid := range table.bb8 {
		droid.Sleep()
	}
}

func cleanup(table droidTable, wg *sync.WaitGroup) {
	sleepAllDroids(table)
	time.Sleep(3 * time.Second)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	table := NewDroidTable()
	argsWithoutProg := os.Args[1:]

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		for _ = range argsWithoutProg {
			cleanup(table, &wg)
		}
	}()

	for _, mrbFile := range argsWithoutProg {
		startWorker(mrbFile, table, &wg)
	}

	wg.Wait()
}
