TAG=v0.1.2

.PHONY: vendor

protoc:
	protoc --go_out=plugins=grpc:./proto *.proto

go-build:
	go build -o exe -mod vendor .

vendor:
	go mod vendor

go-run:
	go run .

run-msgencoder:
	go run ./scripts/msg-encoder/

generate-certs:
	openssl req -x509 -nodes -newkey rsa:2048 -days 365 -keyout privkey.pem -out cert.pem -subj "/CN=127.0.0.1"
	openssl  x509 -in cert.pem -out root.crt

docker-build:
	make vendor
	docker build -t soichisumi0/grpc-bidistream-server:$(TAG) -f ./build/Dockerfile .
	docker push soichisumi0/grpc-bidistream-server:$(TAG)
	rm -rf vendor/