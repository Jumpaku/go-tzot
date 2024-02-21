FROM golang:1.22.0-alpine3.19

RUN apk update && apk add curl jq make