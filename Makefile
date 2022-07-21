BIN=/usr/local/bilibiliscript/bin

init:
	@go build -o ./bs

.PHONY: install
install: init
	@sudo mkdir -p $(BIN) && sudo cp ./bs $(BIN)
	@echo "success!\nmust do it: need write '$(BIN)' to environment variable PATH"
#	@export PATH='$(BIN):$(PATH)'
