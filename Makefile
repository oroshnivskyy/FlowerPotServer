export GOPATH=$(CURDIR)

DEB_VERSION=1


all: gpm build

gpm:
	gpm install

build:
	go build -o=./server ./src

rethink-start:
	docker run --name some-rethink -v "$(CURDIR):/data" -p 8081:8080 -p 28015:28015 -d rethinkdb
