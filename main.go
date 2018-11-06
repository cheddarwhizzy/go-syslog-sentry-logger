package main

import (
	"os"

	"github.com/getsentry/raven-go"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	raven.SetDSN(os.Getenv("SENTRY_DSN"))
	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenTCP("0.0.0.0:10514")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			// fmt.Println(logParts)
			if str, ok := logParts["message"].(string); ok {
				raven.CaptureMessage(str, nil, nil)
			}
		}
	}(channel)

	server.Wait()
}
