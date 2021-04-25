pkgname="bt"

## build: compile bt executable in current directory from source
build:
	go mod tidy
	go build -o "${pkgname}" main.go
	sudo mkdir -p /etc/bash_completion.d
	./"${pkgname}" --completion bash | sudo tee /etc/bash_completion.d/"${pkgname}" > /dev/null
	sudo cp ./"${pkgname}" "${GOROOT}"/bin/"${pkgname}"

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
