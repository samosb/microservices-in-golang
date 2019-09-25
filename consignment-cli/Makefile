build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .
run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli

#build:
#	GOOS=linux GOARCH=amd64 go build
#	docker build -t micro-cli-consignment .
#
#run:
#	docker run micro-cli-consignment
#