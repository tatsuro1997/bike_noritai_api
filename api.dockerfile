FROM golang:1.20

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
