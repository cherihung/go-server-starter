include .env
export

start:
	export ENV="dev" && go run main.go

clean: 
	rm -f ./go-server-starter

build:
	go build -a -installsuffix -o ${MODULE_NAME}

vendorupdate:
	go mod vendor

remove-builder-image:
	docker image prune --filter label=stage=builder -f

docker-dev:
	docker-compose -f docker-compose-dev.yml up -d --build && make remove-builder-image

docker:
	docker-compose up --build -d && make remove-builder-image