PROJNAME := $(notdir $(PWD))
PROJPATH := "github.com/$(GITHUB)/$(PROJNAME)"

all:
	echo $(PROJPATH)
