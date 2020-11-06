# Compile stage
FROM golang:1.15-alpine as build-env

ADD . /go/src/word_games
WORKDIR /go/src/word_games

RUN go build -o /app ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /
COPY --from=build-env /app /

ENTRYPOINT ["/app"]
CMD [""]
