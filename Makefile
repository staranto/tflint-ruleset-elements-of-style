default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	cp ./tflint-ruleset-elements-of-style /tmp
	mv ./tflint-ruleset-elements-of-style ~/.tflint.d/plugins
