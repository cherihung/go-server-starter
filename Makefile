include .env
export

start:
	export ENV="prod" && go run main.go

clean: 
	rm -f ./go-server-starter

dev-build:
	GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -a installsuffix -o ${MODULE_NAME}

vendorupdate:
	go mod vendor
