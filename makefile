PWD=$(shell pwd)

all: build

css:
	compass compile

build: css
	go build

run: build
	$(PWD)/keepreading
