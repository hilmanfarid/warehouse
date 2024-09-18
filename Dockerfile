FROM golang:alpine

WORKDIR /app

ADD . /app/

RUN go build -o ./out/warehouse-project .

EXPOSE 8000

ENTRYPOINT ["./out/warehouse-project"]