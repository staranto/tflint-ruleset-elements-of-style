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
	@if [ -z "$(VERSION)" ]; then echo "Usage: make release VERSION=x.y.z"; exit 1; fi
	@if ! echo "$(VERSION)" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$$'; then \
		echo "Error: VERSION must be a valid semantic version (e.g. 0.2.0) without leading 'v'. Got: $(VERSION)"; \
		exit 1; \
	fi
	sed -i 's/version = "[0-9]\+\.[0-9]\+\.[0-9]\+"/version = "$(VERSION)"/' README.md
 	git add README.md
 	git commit -m "chore: bump version to $(VERSION)"
 	git tag v$(VERSION)
	@echo "Successfully bumped to $(VERSION) and created tag v$(VERSION)."
	@echo "Run 'git push origin main --tags' to publish."
