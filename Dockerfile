FROM golang

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go test -c -o external_test

EXPOSE 8080

ENTRYPOINT exec /app/external_test