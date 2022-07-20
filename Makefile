user :=	$(shell whoami)

ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
	ifneq ($(GOPATH),)
		BIN=$(GOPATH)/bin
	endif
endif

init:
	@cd bin && go build -o ../bs && cd - &> /dev/null

.PHONY: install
install: init
ifeq ($(user),root)
#root, install for all user
	@cp ./bs /usr/bin
	@echo "Install success, bin to path: /usr/bin"
else
#!root, install for current user
	$(shell if [ -z $(BIN) ]; then read -p "Please select installdir: " REPLY; mkdir -p $${REPLY};\
	cp ./bs $${REPLY}/;else mkdir -p $(BIN);\
	cp ./bs $(BIN); fi)
	@echo "Install success, bin to path: $(BIN)"
endif
