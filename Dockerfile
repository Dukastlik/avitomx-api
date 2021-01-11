FROM golang:1.15.6-buster

ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .
CMD make

EXPOSE 8080
RUN ls

ENTRYPOINT ["./apiserver"]