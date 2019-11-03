start:
	export ENV="prod" && go run main.go

clean: 
	rm -f ./go-server-starter

build:
	make clean && go build

vendorupdate:
	go mod vendor

devcerts:
	cd _certs && 
	openssl genrsa -out server.key 2048 && 
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650