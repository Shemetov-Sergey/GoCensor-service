FROM golang:alpine as builder

WORKDIR /GoCensor-service

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine

WORKDIR /GoCensor-service

COPY --from=builder /GoCensor-service/main /GoCensor-service/main
COPY --from=builder /GoCensor-service/pkg/config/envs/*.env /GoCensor-service/

RUN chmod +x /GoCensor-service/main

CMD ["./main"]