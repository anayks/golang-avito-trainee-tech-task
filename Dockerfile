FROM golang:latest

ENV GOPATH=/

RUN apt-get update
RUN apt-get -y install postgresql-client

COPY ./ /go/src/golang-avito-tech-test
WORKDIR /go/src/golang-avito-tech-test
RUN chmod +x ./wait-for-postgres.sh
RUN go mod download
RUN go build -o /go/src/golang-avito-tech-test/main .

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate
CMD ["./main"]
