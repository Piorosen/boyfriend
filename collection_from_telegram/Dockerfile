FROM golang:1.22.4-alpine3.20 AS builder
RUN apk add git g++ gcc

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .

FROM scratch
COPY --from=builder /dist/main .
ENTRYPOINT ["/main"]
