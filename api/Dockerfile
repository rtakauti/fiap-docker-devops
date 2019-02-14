FROM golang:1.11.5-alpine3.8 as build-stage

RUN apk add --update git

WORKDIR /go/src/api
COPY datamodels/ ./datamodels/
COPY repositories/ ./repositories/
COPY services/ ./services/
COPY shared/ ./shared/
COPY main.go ./

RUN go get -d -v ./...
RUN go install

FROM golang:1.11.5-alpine3.8 as production-stage
EXPOSE 80
CMD ["api"]