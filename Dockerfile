FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd ./cmd
COPY pkg ./pkg

RUN go mod download

RUN go build -o myapp ./cmd

CMD ["./myapp"]
