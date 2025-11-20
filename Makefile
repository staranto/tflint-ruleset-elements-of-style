.PHONY: default test build install release

default: build

test:
	go test ./...

build: test
	go build

install: build
	@mkdir -p ~/.tflint.d/plugins
	@mv ./tflint-ruleset-elements-of-style ~/.tflint.d/plugins
	@echo "Successfully installed tflint-ruleset-elements-of-style to ~/.tflint.d/plugins"

release:
	@if [ -z "$(VERSION)" ]; then echo "Usage: make release VERSION=0.x.y"; exit 1; fi
	@sed -i 's/version = "[0-9]\+\.[0-9]\+\.[0-9]\+"/version = "$(VERSION)"/' README.md
	@git add README.md
	@git commit -m "chore: bump version to $(VERSION)"
	@git tag v$(VERSION)
	@echo "Successfully bumped to $(VERSION) and created tag v$(VERSION)."
	@echo "Run 'git push origin main --tags' to publish."
