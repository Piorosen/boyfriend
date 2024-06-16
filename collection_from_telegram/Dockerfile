FROM golang:1.22.4-alpine3.20 AS builder
RUN apk add --no-cache git g++ gcc

ARG VERSION=NULL
ARG DATE=NULL

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main -ldflags "-X 'main.BUILD_DATE_TIME=${DATE}' -X 'main.DEPLOY_VERSION=${VERSION}'" . 
WORKDIR /dist
RUN cp /build/main .

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/Piorosen/boyfriend
LABEL org.opencontainers.image.licenses=MIT
LABEL org.opencontainers.image.authors="JooHyoung Cha"
COPY --from=builder /dist/main .
ENTRYPOINT ["/main"]
