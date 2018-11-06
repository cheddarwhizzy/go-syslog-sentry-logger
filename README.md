## Setup
- On your linux system running rsyslog as a service, add a file called `/etc/rsyslog.d/50-sentry.conf` with one of the following

#### A: To log all messages (all facilities)
```
*.* @@ip-or-hostname-of-this-service:10514;RSYSLOG_SyslogProtocol23Format
```

#### B: To log only error messages
```
local7.err @@ip-or-hostname-of-this-service:10514;RSYSLOG_SyslogProtocol23Format
```

## How to Use
### Run in Docker
```
docker run -d --name sentry-syslogger \
    -p 10514:10514 \
    -e SENTRY_DSN="https://user:pass@sentry.hostname.com/sdfsfsdsdf/7"
    bporter2387/sentry_syslogger:latest
```

### Run in Kubernetes
```
kubectl run sentry-syslogger --image=bporter2387/sentry-syslogger:latest --expose --port=10514 --env SENTRY_DSN="https://user:pass@sentry.hostname.com/sdfsfsdsdf/7"
```

### Run command line
```
SENTRY_DSN=https://user:pass@sentry.hostname.com/sdfsfsdsdf/7 go run main.go
```


### Testing
```
logger -p local7.err "This is a test. you should see it in sentry"
```