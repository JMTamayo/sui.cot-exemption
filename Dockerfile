FROM ubuntu:24.04 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

RUN apt-get update && \
    apt-get install -y golang-go git && \
    rm -rf /var/lib/apt/lists/*

COPY . .
RUN go mod download

RUN go build -buildvcs=false -o main ./app/main.go

WORKDIR /dist

RUN cp /build/main .

FROM ubuntu:24.04

COPY --from=builder /dist/main /main

RUN apt-get update && \
    apt-get install -y --no-install-recommends wget apt-transport-https ca-certificates && \
    wget -qO /etc/apt/trusted.gpg.d/nordvpn_public.asc https://repo.nordvpn.com/gpg/nordvpn_public.asc && \
    echo "deb https://repo.nordvpn.com/deb/nordvpn/debian stable main" > /etc/apt/sources.list.d/nordvpn.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends nordvpn && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

EXPOSE 8000

CMD ["/bin/sh", "-c", "/etc/init.d/nordvpn start && sleep 5 && ./main"]
