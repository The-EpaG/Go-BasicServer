FROM golang:1.23.4 AS base

WORKDIR /app

FROM base AS copy

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY makefile .
COPY main.go .
COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal
COPY config/config.go ./config/config.go
COPY config/config.yaml ./config/config.yaml

FROM copy AS test

COPY test ./test
COPY config/test.yaml ./config/test.yaml

RUN make test

FROM copy AS lint

RUN make lint

FROM copy AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN make build

FROM alpine:latest AS final

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=build /app/main /app
COPY --from=build /app/config/config.yaml /app/config/config.yaml

RUN addgroup -S user && adduser -S user -G user

RUN chown -R user:user /app

USER user

EXPOSE 8080

ENV GIN_MODE=release

CMD [ "/app/main", "serve" ]

