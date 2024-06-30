.PHONY: run, clear, connect
clear:
	clear
	
run: clear
	@go run ./cmd .

connect: clear
	./scripts/connect.sh

server: clear
	@go run ./server/ .
