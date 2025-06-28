FROM golang:1.24-alpine AS build
RUN apk add --no-cache build-base sqlite-dev
WORKDIR /src
COPY go.mod ./
RUN go mod download && go mod tidy

COPY . .

RUN go build -o pack-calculator ./cmd/server
FROM alpine:3.20
RUN apk add --no-cache ca-certificates sqlite-libs

WORKDIR /srv
COPY --from=build /src/pack-calculator .
COPY ui ./ui

ENV PACK_CALC_PORT=:8081
EXPOSE 8081
ENTRYPOINT ["./pack-calculator"]