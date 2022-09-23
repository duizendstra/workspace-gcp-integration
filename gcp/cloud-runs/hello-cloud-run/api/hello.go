package api

import (
	// "context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"duizendstra.com/gcp"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	event_http "github.com/cloudevents/sdk-go/v2/protocol/http"
)

func GetReceive() func(event cloudevents.Event) protocol.Result {
	return func(event cloudevents.Event) protocol.Result {
		// check if we have the correct event
		if event.Type() != "google.cloud.pubsub.topic.v1.messagePublished" {
			log.Println(gcp.Entry{
				Severity:  "WARNING",
				Message:   fmt.Sprintf("Recieved an unsupported event type: %s", event.Type()),
				Component: gcp.LogComponent,
			})
			return event_http.NewResult(200, "%w", protocol.ResultACK)
		}

		obj := &PubSubEvent{}
		if err := event.DataAs(obj); err != nil {
			log.Print(
				gcp.Entry{
					Severity:  "WARNING",
					Message:   fmt.Sprintf("Unable to decode PubSub data: %s\n", err.Error()),
					Component: gcp.LogComponent,
				})
			return event_http.NewResult(200, "%w", protocol.ResultACK)
		}

		decodedMessage, err := base64.StdEncoding.DecodeString(obj.Message.Data)

		if err != nil {
			log.Println(gcp.Entry{
				Severity:  "WARNING",
				Message:   fmt.Sprintf("Unable to decode PubSub message: %s\n", err.Error()),
				Component: gcp.LogComponent,
			})
			return event_http.NewResult(200, "%w", protocol.ResultACK)
		}

		log.Println(gcp.Entry{
			Severity:  "DEBUG",
			Message:   fmt.Sprintf("Received event of type %s. Event data: %s", event.Type(), decodedMessage),
			Component: gcp.LogComponent,
		})

		var logEntry LogEntry
		if err = json.Unmarshal(decodedMessage, &logEntry); err != nil {
			log.Println(gcp.Entry{
				Severity: "ERROR",
				Message:  fmt.Sprintf("Error parsing application/json: %v", err),
			})
			return event_http.NewResult(200, "%w", protocol.ResultACK)
		}

		if logEntry.ProtoPayload.MethodName != "google.admin.AdminService.moveUserToOrgUnit" {
			log.Println(gcp.Entry{
				Severity:  "ERROR",
				Message:   fmt.Sprintf("Missing or invalid payload %v", logEntry.ProtoPayload),
				Component: gcp.LogComponent,
			})
			return event_http.NewResult(200, "%w", protocol.ResultACK)
		}

		log.Println(gcp.Entry{
			Severity: "INFO",
			Message:  "Hello cloud run!",
		})

		return event_http.NewResult(200, "%w", protocol.ResultACK)
	}
}
