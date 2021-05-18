GC=go

PROJECT=palacinke

MODULE=github.com/juanvillacortac/palacinke

MAIN_PKG=$(MODULE)/cmd/palacinke

OUTPUT_DIR=.

ifeq ($(GOOS), windows)
	EXE := .exe
endif

ifdef COMSPEC
	EXE := .exe
endif

OUTPUT=$(OUTPUT_DIR)/$(PROJECT)$(EXE)

.PHONY: run build fmt clean gen test

all: build

run:
	@$(GC) run $(MAIN_PKG) $(args)

build:
	@echo "Compiling to $(OUTPUT)"
	@$(GC) build -o $(OUTPUT) $(MAIN_PKG)
	@echo "Done!"

fmt:
	@$(GC) fmt ./...

clean:
	@rm $(OUTPUT_DIR) -r

gen:
	@echo "Generating all..."
	@$(GC) generate ./...
	@echo "Done!"

test:
	@$(GC) test -count=1 ./test/...
