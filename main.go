package main

import (
	"fmt"
	"os"

	"github.com/getsentry/raven-go"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	raven.SetDSN(os.Getenv("SENTRY_DSN"))
	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Using Sentry DSN:", os.Getenv("SENTRY_DSN"))
	}
	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenTCP("0.0.0.0:10514")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			info := map[string]string{
				"message":  logParts["message"].(string),
				"hostname": logParts["hostname"].(string),
			}
			if os.Getenv("DEBUG") == "true" {
				fmt.Println(logParts)
				fmt.Println("Sending", info)
			}
			if str, ok := logParts["message"].(string); ok {
				raven.CaptureMessage(str, info, nil)
			}
		}
	}(channel)

	server.Wait()
}
