all: dep mruby build

dep:
	dep ensure

mruby:
	cd vendor/github.com/mitchellh/go-mruby; MRUBY_CONFIG=../../../../../../build_config.rb make libmruby.a

build:
	go build -o build/mrgoboto main.go

clean:
	rm -rf build/ vendor/

.PHONY: all dep mruby build clean
