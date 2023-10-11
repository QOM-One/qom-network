FROM golang:1.21.3-bullseye AS build-env

WORKDIR /go/src/github.com/qom-one/qomapp

RUN apt-get update -y
RUN apt-get install git -y

COPY . .

RUN make build

FROM golang:1.21.3-bullseye

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/qom-one/qomapp/build/qomd /usr/bin/qomd

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["qomd"]
