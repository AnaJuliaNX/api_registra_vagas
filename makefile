.PHONY: default run bild test docs clean

#variables
APP_NAME=registraVagas

# Tasks
default: run-dois

run: 
	@go run main.go
run-dois:
	@swag init 
	@go run main.go
bild: 
	@go build -o $(APP_NAME) main.go
test: 
	@go test ./ ...
docs: 
	@swag init 
clean:
	@rm -f $(APP_NAME)