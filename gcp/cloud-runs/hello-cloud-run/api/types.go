package api

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"time"
)

type LogEntry struct {
	InsertID     string `json:"insertId"`
	LogName      string `json:"logName"`
	ProtoPayload struct {
		Type               string `json:"@type"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		AuthorizationInfo []struct {
			Granted    bool   `json:"granted"`
			Permission string `json:"permission"`
			Resource   string `json:"resource"`
		} `json:"authorizationInfo"`
		Metadata struct {
			Type            string `json:"@type"`
			Group           string `json:"group"`
			MembershipDelta struct {
				Member     string `json:"member"`
				RoleDeltas []struct {
					Action string `json:"action"`
					Role   string `json:"role"`
				} `json:"roleDeltas"`
			} `json:"membershipDelta"`
		} `json:"metadata"`
		MethodName      string `json:"methodName"`
		RequestMetadata struct {
			CallerIP                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
		} `json:"requestMetadata"`
		ResourceName string `json:"resourceName"`
		ServiceName  string `json:"serviceName"`
	} `json:"protoPayload"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
	Resource         struct {
		Labels struct {
			Method  string `json:"method"`
			Service string `json:"service"`
		} `json:"labels"`
		Type string `json:"type"`
	} `json:"resource"`
	Severity  string    `json:"severity"`
	Timestamp time.Time `json:"timestamp"`
}

type LoggedEvent struct {
	Severity  string            `json:"severity"`
	EventType string            `json:"eventType"`
	Message   string            `json:"message"`
	Event     cloudevents.Event `json:"event"`
}
type PubSubMessage struct {
	Data string `json:"data"`
}
type PubSubEvent struct {
	Message PubSubMessage `json:"message"`
}
