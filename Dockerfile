FROM golang:latest

WORKDIR /go/src/cheddarwhizzy/sentry_syslogger/
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    go build -o /app .

EXPOSE 10514

CMD ["/app"]