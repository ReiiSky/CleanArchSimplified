PROJECT_NAME=CLEAN_ARCH
PROJECT_VERSION=1.0
EXPOSED_PORT=8080

run-dev: 
	go run main.go

build: 
	go build -o server main.go

build-win:
	GOOS=windows GOARCH=386 go build -o server main.go

docker-img:
	docker build --tag $(PROJECT_NAME):$(PROJECT_VERSION) .

docker-run:
	docker run -d -p $(EXPOSED_PORT):6007 $(PROJECT_NAME):$(PROJECT_VERSION)

docker: docker-img docker-run
