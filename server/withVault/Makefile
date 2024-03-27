BINARY_NAME=myapp
VAULT_TOKEN="myroot"
VAULT_PATH="/secret"
DATABASE_USER_SECURE="/database:username"
DATABASE_PASSWORD_SECURE="/database:password"
VAULT_ADDRESS="http://0.0.0.0:8300"
LOCAL_BIN:=$(CURDIR)/bin


## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ./bin/${BINARY_NAME} ./cmd
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env VAULT_TOKEN=${VAULT_TOKEN} VAULT_PATH=${VAULT_PATH} VAULT_ADDRESS=${VAULT_ADDRESS} DATABASE_USER_SECURE=${DATABASE_USER_SECURE} DATABASE_PASSWORD_SECURE=${DATABASE_PASSWORD_SECURE}  ./bin/${BINARY_NAME} &
	@echo "Started!"

start: run

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"
	@echo "Stopped!"

restart: stop start