jq=$(shell which jq)
yq=$(shell which yq)

all:
	@# Check if jq installed
ifeq ($(strip $(jq)),)
	@echo "jq not found. Consider doing sudo apt-get install jq"
else
	@echo "Found jq"
endif
	@# Check if yq installed
ifeq ($(strip $(yq)),)
	@echo "yq not found. Consider doing pip3 install yq"
else
	@echo "Found yq"
endif
	@echo "Done"

install:
	sudo cp src/Dockerfile/ga /usr/bin/

uninstall:
	sudo rm /usr/bin/ga

