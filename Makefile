all: clean build run

build:
	go build -o ./bin/web .
run:
	./bin/web
clean:
	rm -rf ./bin
sqlc:
	sqlc generate -f ./database/sqlc.yaml
