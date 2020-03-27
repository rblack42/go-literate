PROJNAME := $(notdir $(PWD))
PROJPATH := "github.com/$(GITHUB)/$(PROJNAME)"

run:
	go run .

init:
	go mod init $(GOPATH)

debug:
	echo $(PROJPATH)
