start:
	export ENV="prod" && go run main.go

clean: 
	rm -f ./go-server-starter

build:
	make clean && go build

vendorupdate:
	go mod vendor