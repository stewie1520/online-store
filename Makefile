all: clean build run

debug: clean build dlv

build:
	go build -o ./bin/web .
run:
	./bin/web
clean:
	rm -rf ./bin
gen-sqlc:
	sqlc generate -f ./database/sqlc.yaml
gen-wire:
	wire ./...
dlv:
	dlv exec ./bin/web --listen=:8181 --headless=true --api-version=2 --accept-multiclient
