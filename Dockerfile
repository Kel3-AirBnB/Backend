FROM golang:1.22.2-alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

COPY ./ /app

RUN go mod tidy

RUN go build -o airbnbv1.2

# CMD [ "./airbnbv1.2" ]
# Tambahkan langkah debug untuk memeriksa apakah environment variable ada
CMD echo "Starting app with env variables:" && \
    echo "DBUSER=${DBUSER}" && \
    echo "S3BUCKETNAME=${S3BUCKETNAME}" && \
    ./airbnbv1.2