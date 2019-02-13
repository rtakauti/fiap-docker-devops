FROM goland:1.11.5-alpine3.8 as build-stage

WORKDIR /go/src/app
COPY web .
COPY repositories .
COPY services .
COPY shared .
COPY web .

RUN go get -d -v ./...
RUN go install -v ./...

FROM goland:1.11.5-alpine3.8 as production-stage
CMD ["app"]