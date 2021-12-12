all: clean build

build:
	go build -o ./bin/web .
	./bin/web
clean:
	rm -rf ./bin