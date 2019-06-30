FROM golang:1.8 as build-env

WORKDIR /go/src/app

COPY . .

RUN go build -o Main src/Main.go

FROM golang:1.8

WORKDIR /go/src/app

COPY --from=build-env /go/src/app/Main /go/bin

CMD ["Main"]