.PHONY: all test cover gen get
COVER_FILE=./.coverage.txt
MOCKS=./mocks
BIN=bin/

all: get build test

default: all

get:
	go install github.com/golang/mock/mockgen@v1.6.0 && go get -u github.com/dave/courtney && go get ./...

gen: get
	go generate ./...

build: get
	go build -o $(BIN) ./...

test: gen
	go install github.com/dave/courtney && courtney -o $(COVER_FILE) ./pkg/... && go tool cover -func $(COVER_FILE)

clean:
	$(GOCLEAN)
	rm -rf $(BIN) || true && rm $(COVER_FILE) || true && rm -rf $(MOCKS)

install:
	go install ./...

cover: test
	go tool cover -func $(COVER_FILE)

docker-build:
	(cd docker/builder-ubuntu && make run)