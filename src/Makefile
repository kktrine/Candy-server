
generate:
	@swagger generate server -f ./swagger/swagger.yaml
client:
	@go build -o candy-client cmd/client/client.go
minica:
	@if [ ! -f minica.pem ]; then \
		@git clone https://github.com/jsha/minica.git; \
		@chmod +x ./minica/minica; \
		@cd ./minica && go build -o minicagen ./main.go; \
		@./minica/minicagen --domains candy.tld; \
		@rm -f minicagen; \
		@rm -rf ./minica; \
	fi

server:
	@go run cmd/candy-server-server/main.go --tls-certificate=./candy.tld/cert.pem --tls-key=./candy.tld/key.pem
run:
	@make generate
	@make minica
	@make client
	@make server
