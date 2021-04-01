TINYGOFLAGS = -target wasm -wasm-abi generic

SRC = ./src
WASM_DEBUG = build/soralet.wasm
WASM_RELEASE = build/soralet-optimized.wasm

help:
	@echo "usage: make <\033[36mtargets\033[0m>"
	@echo
	@echo "available targets:"
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: $(SRC) ## Build debug WASM
	tinygo build -o $(WASM_DEBUG) $(TINYGOFLAGS) $(SRC)

release: $(SRC) ## Build release WASM
	tinygo build -o $(WASM_RELEASE) $(TINYGOFLAGS) -no-debug $(SRC)

upload:
	soracom soralets upload --soralet-id my-soralet --content-type application/octet-stream --body @${WASM_RELEASE}

test:
	soracom soralets exec --soralet-id my-soralet --version 1 --direction uplink --content-type application/json --body @test.json

clean: ## Remove WASM(s) under build/
	rm -fr build/*.wasm

.PHONY: help build release clean
