FROM golang:latest

WORKDIR /go/src/cheddarwhizzy/sentry_syslogger/

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . .

RUN dep ensure && \
    go build -o /app .

EXPOSE 10514

CMD ["/app"]