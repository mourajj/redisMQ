FROM golang:alpine

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/mourajj/redisMQ/main
RUN cd /build && git clone https://github.com/mourajj/redisMQ.git

RUN cd /build/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/RedisMQ/main" ]