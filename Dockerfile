FROM golang:1.26.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /tf-module-graph ./cmd/tf-module-graph

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /tf-module-graph /usr/local/bin/tf-module-graph
ENTRYPOINT ["/usr/local/bin/tf-module-graph"]