FROM golang:alpine as builder

RUN apk update && apk add --no-cache 'git=~2'

WORKDIR /src/packages/app

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main

FROM alpine:3

WORKDIR /

COPY --from=builder /go/main /go/main
COPY . /go/public

EXPOSE 8080

WORKDIR /go

ENTRYPOINT [ "./main" ]