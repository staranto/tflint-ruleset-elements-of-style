default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	cp ./tflint-ruleset-type-echo /tmp
	mv ./tflint-ruleset-type-echo ~/.tflint.d/plugins
