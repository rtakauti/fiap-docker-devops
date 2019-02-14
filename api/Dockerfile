FROM golang:1.11.5-alpine3.8 as build-stage

RUN apk add --update git

WORKDIR /go/src/app
COPY *.go ./

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80
CMD ["app"]