FROM golang:1.24-alpine AS builder


RUN apk add --no-cache git


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o todoapp ./cmd


FROM alpine:latest


RUN apk add --no-cache tzdata


WORKDIR /app


COPY --from=builder /app/todoapp .
COPY templates ./templates


EXPOSE 8080


CMD ["./todoapp"]