FROM golang:1.22.2-alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod tidy

COPY ./ /app

RUN go build -o airbnbv1

CMD [ "./airbnbv1" ]