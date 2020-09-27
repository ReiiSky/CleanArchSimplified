FROM golang:1.15.0-alpine3.12

RUN mkdir /app
WORKDIR /app
COPY . /app
EXPOSE 6007
CMD [ "go", "run", "main.go" ]