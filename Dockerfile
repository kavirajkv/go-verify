FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /goverify


FROM alpine:latest

WORKDIR /
COPY --from=builder /goverify /goverify

EXPOSE 8080

ENTRYPOINT ["/goverify"]