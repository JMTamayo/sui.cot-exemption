FROM ubuntu:24.04 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Install Go and build dependencies
RUN apt-get update && \
    apt-get install -y golang-go git && \
    rm -rf /var/lib/apt/lists/*

COPY . .
RUN go mod download

RUN go build -buildvcs=false -o main ./app/main.go

WORKDIR /dist

RUN cp /build/main .

FROM ubuntu:24.04

# Install runtime dependencies if needed
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /dist/main /main

EXPOSE 8000

ENTRYPOINT ["/main"]