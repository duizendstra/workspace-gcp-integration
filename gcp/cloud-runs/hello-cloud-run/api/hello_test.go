package api

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
	http "github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/cloudevents/sdk-go/v2/types"
	"github.com/google/go-cmp/cmp"
	"net/url"
	"testing"
)

func MinEventContextV01() *event.EventContextV03 {
	sourceUrl, _ := url.Parse("http://example.com/source")
	source := &types.URIRef{URL: *sourceUrl}

	return event.EventContextV03{
		Type:   "com.example.simple",
		Source: *source,
		ID:     "ABC-123",
	}.AsV03()
}

func MinEventContextV02() *event.EventContextV03 {
	sourceUrl, _ := url.Parse("http://example.com/source")
	source := &types.URIRef{URL: *sourceUrl}

	return event.EventContextV03{
		Type:   "google.cloud.pubsub.topic.v1.messagePublished",
		Source: *source,
		ID:     "ABC-123",
	}.AsV03()
}

func MinEventContextV03() *event.EventContextV03 {
	sourceUrl, _ := url.Parse("http://example.com/source")
	source := &types.URIRef{URL: *sourceUrl}

	return event.EventContextV03{
		Type:                "google.cloud.pubsub.topic.v1.messagePublished",
		Source:              *source,
		ID:                  "ABC-123",
		DataContentType:     event.StringOfApplicationJSON(),
		DataContentEncoding: event.StringOfBase64(),
	}.AsV03()
}

func AddMoveUserToOrgUnitPayload() string {
	data := `{"protoPayload":{"@type":"type.googleapis.com/google.cloud.audit.AuditLog","authenticationInfo":{"principalEmail":"user1@example.com"},"requestMetadata":{"callerIp":"127.0.0.1","requestAttributes":{},"destinationAttributes":{}},"serviceName":"admin.googleapis.com","methodName":"google.admin.AdminService.moveUserToOrgUnit","resourceName":"organizations/1234567890/userSettings","metadata":{"@type":"type.googleapis.com/ccc_hosted_reporting.ActivityProto","activityId":{"timeUsec":"1663476083593000","uniqQualifier":"-8299540676987354861"},"event":[{"parameter":[{"type":"TYPE_STRING","label":"LABEL_OPTIONAL","value":"user2@example.com","name":"USER_EMAIL"},{"name":"ORG_UNIT_NAME","value":"/Admins","type":"TYPE_STRING","label":"LABEL_OPTIONAL"},{"type":"TYPE_STRING","name":"NEW_VALUE","label":"LABEL_OPTIONAL","value":"/"}],"eventType":"USER_SETTINGS","eventName":"MOVE_USER_TO_ORG_UNIT"}]}},"insertId":"fjwwoaeb955f","resource":{"type":"audited_resource","labels":{"service":"admin.googleapis.com","method":"google.admin.AdminService.moveUserToOrgUnit"}},"timestamp":"2022-09-18T04:41:23.593Z","severity":"NOTICE","logName":"organizations/506102632746/logs/cloudaudit.googleapis.com%2Factivity","receiveTimestamp":"2022-09-18T04:41:24.500756352Z"}`
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return fmt.Sprintf(`{"message": {"data":"%s"}}`, sEnc)
}

func TestGroupUpdateMembership(t *testing.T) {
	eventTest1 := event.Event{
		Context: MinEventContextV01(),
	}

	eventTest2 := event.Event{
		Context:     MinEventContextV02(),
		DataEncoded: []byte("Here is a string...."),
	}

	eventTest3 := event.Event{
		Context: MinEventContextV03(),
	}

	eventTest3.SetData("application/json", []byte(AddMoveUserToOrgUnitPayload()))

	type args struct {
		r    event.Event
		want protocol.Result
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "EventPost", args: args{
			r:    eventTest1,
			want: http.NewResult(200, "%w", protocol.ResultACK),
		}},
		{name: "EventPost", args: args{
			r:    eventTest2,
			want: http.NewResult(200, "%w", protocol.ResultACK),
		}},
		{name: "EventPost", args: args{
			r:    eventTest3,
			want: http.NewResult(200, "%w", protocol.ResultACK),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receive := GetReceive()
			got := receive(tt.args.r)

			if !cmp.Equal(tt.args.want.Error(), got.Error()) {
				t.Errorf("want %v got %v", tt.args.want.Error(), got.Error())
			}
		})
	}
}
