BINARY="im"


# colors compatible setting
CRED:=$(shell tput setaf 1 2>/dev/null)
CGREEN:=$(shell tput setaf 2 2>/dev/null)
CYELLOW:=$(shell tput setaf 3 2>/dev/null)
CEND:=$(shell tput sgr0 2>/dev/null)


.PHONY: all
all: build

# build
.PHONY: build
build:
	@echo "$(CGREEN)Building for linux...$(CEND)"
	[ -d "bin" ] || sudo mkdir bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${BINARY} cmd/main.go
	@echo "$(CGREEN)Build Success!$(CEND)"

# clean
.PHONY: clean
clean:
	@echo "$(CGREEN)Cleanup...$(CEND)"
	go clean
	@rm bin/${BINARY}
	@echo "rm bin/${BINARY}"