.PHONY: \
	dep \
	install \
	build \
	vet \
	test

dep:
	go mod download

build: main.go
	go build -o grappet $<

install: main.go
	go install

test:
	go test ./...

vet:
	go vet

setup:
	./setup.sh
	go build -o grappet $<
