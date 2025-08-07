FROM golang:1.24.5-alpine

RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o hospital-api ./cmd/main.go

EXPOSE 8002

CMD ["./hospital-api"]
