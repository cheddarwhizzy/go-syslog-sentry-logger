
add a file called /etc/rsyslog.d/50-sentry.conf

# *.* will log all syslog messages
# change to something like 'local7.err' for only error messages
*.*       @@ip-or-hostname-of-command-below:10514;RSYSLOG_SyslogProtocol23Format


# Running this will spin up a TCP server listening on 10514
SENTRY_DSN=https://user:pass@sentry.hostname.com/sdfsfsdsdf/7 go run main.go
