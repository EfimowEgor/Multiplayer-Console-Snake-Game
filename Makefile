.PHONY: run, clear
clear:
	clear
run: clear
	@go run ./cmd .

server: clear
	@go run ./server/ .
