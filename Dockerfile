FROM golang

RUN mkdir /usr/src/app

ADD . /usr/src/app

WORKDIR /usr/src/app

RUN go mod download

RUN go build -o main .


CMD ["/usr/src/app/main"]