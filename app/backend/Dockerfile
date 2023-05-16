# syntax=docker/dockerfile:1
FROM golang:1.17-alpine as builder
WORKDIR /backend
RUN apk add --no-cache git
COPY . .
RUN go mod download
RUN go mod tidy
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /background/main .
CMD ["./main"]  
