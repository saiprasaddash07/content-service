FROM golang:latest

WORKDIR /content-service.git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./content-service.git"]