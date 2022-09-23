package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "duizendstra.com/api"
	"duizendstra.com/gcp"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func init() {
	// Disable log prefixes such as the default timestamp.
	// Prefix text prevents the message from being parsed as JSON.
	// A timestamp is added when shipping logs to Cloud Logging.
	log.SetFlags(0)
	gcp.LogComponent = "hello-cloud-run"
}

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Println(gcp.Entry{
			Severity:  "DEBUG",
			Message:   fmt.Sprintf("Failed to create the cloudevents HTTP client: %v", err),
			Component: gcp.LogComponent,
		})
	}

	result := c.StartReceiver(context.Background(), api.GetReceive())
	log.Println(gcp.Entry{
		Severity:  "DEBUG",
		Message:   fmt.Sprintf("Failed to start the receiver: %v", result),
		Component: gcp.LogComponent,
	})

	os.Exit(1)
}
