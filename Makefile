GIT_COMMIT=$(shell git describe --always)

all: build
default: build

build:
	go build -o wechat-channel

clean:
	rm wechat-channel