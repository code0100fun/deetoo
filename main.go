package main

import (
    "io/ioutil"
    "fmt"
    "sync"
    "github.com/mitchellh/go-mruby"
)

func apiSendMessage(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
    args := m.GetArgs()
    channel := args[0].String()
    message := args[1].Hash()
    fmt.Printf("Channel: %s, Message: %s\n", channel, message)
    return nil, nil
}

func startWorker(file string, wg *sync.WaitGroup) {
    wg.Add(1)
    go func() {
        defer wg.Done()
        mrb := mruby.NewMrb()
        defer mrb.Close()

        class := mrb.DefineClass("Go", nil)
        class.DefineClassMethod("send_message", apiSendMessage, mruby.ArgsReq(2))

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
    startWorker("main.mrb", &wg)
    startWorker("main.mrb", &wg)
    startWorker("main.mrb", &wg)
    startWorker("main.mrb", &wg)
    startWorker("main.mrb", &wg)
    startWorker("main.mrb", &wg)
    startWorker("main.mrb", &wg)
    wg.Wait()
}
