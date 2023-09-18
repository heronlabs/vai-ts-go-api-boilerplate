download: 
	go mod download

build: 
	rm -rf bin && go build -o ./bin/ ./src/*.go

lint: 
	golangci-lint run ./src

script-run: 
	go run ./src/main.go

docker-build: 
	docker build -t go-api:latest . -f Dockerfile.local

docker-up-script: 
	docker-compose up -d go-api --build --force-recreate

docker-down: 
	docker-compose down --rmi all