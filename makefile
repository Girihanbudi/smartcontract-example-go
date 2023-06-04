# Run a dependencies injection using google wire
dependencies: 
	cd ./internal/app && wire

# Generate application's swagger documentation
documentation:
	swag init -g ./cmd/app/main.go -o ./docs

# Build all smart contract sdk
contractsdk:
	solcjs --optimize --abi --bin ./internal/app/filesystem/filesystem.sol -o build
	abigen --abi=./build/internal_app_filesystem_filesystem_sol_FileSystem.abi --bin=./build/internal_app_filesystem_filesystem_sol_FileSystem.bin --pkg=filesystem --out=./internal/app/filesystem/smartcontract.go

# Create all environments needed in docker compose
environments:
	docker-compose up -d

# Deploy contract to ethereum server
contractdeployment:
	go run ./cmd/contract/main.go

# Run application
run:
	go run ./cmd/app/main.go

# Create application image
image:
	docker build -t smartm2m-service .

# Run a container from created image 
container:
	docker run -d -it -p 8081:8081 --name=smartm2m-service smartm2m-service

# Generate application image and run the container
rundocker: image container