FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /scraper

COPY . /scraper/

RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build scraper

EXPOSE $PORT

ENTRYPOINT ["./scraper"]