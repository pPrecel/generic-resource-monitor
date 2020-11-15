FROM golang:1.15-alpine as builder

ENV BASE_APP_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR ${BASE_APP_DIR}

# Copy the Go Modules manifests
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Copy files
COPY . ${BASE_APP_DIR}/

# Build
RUN go build -a -o main cmd/main.go \
    && mv main /app/main

FROM scratch

COPY --from=builder /app /app

ENTRYPOINT ["/app/main"]
