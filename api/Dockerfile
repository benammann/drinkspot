FROM golang:latest

ENV CGO_ENABLED 0

WORKDIR /application

COPY ./ /application

RUN cd src && go mod download

RUN go get -u github.com/cosmtrek/air

ENTRYPOINT air