package datadog_test

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

var testEvent = datadog.Event{
	Title: "test event from go client",
	Text:  "example text",
	Tags: &[]string{
		"test",
		"client:go",
	},
	Priority:       datadog.PtrString("normal"),
	SourceTypeName: datadog.PtrString("datadog-api-client-go"),
}

func TestEventLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// Create event
	eventResponse, httpresp, err := TESTAPICLIENT.EventsApi.CreateEvent(TESTAUTH).Event(testEvent).Execute()
	if err != nil {
		t.Fatalf("Error creating Event %v: Response %s: %v", testEvent, err.(datadog.GenericOpenAPIError).Body(), err)
	}
	// defer deleteEvent(eventResponse.GetEvent().GetId())
	assert.Equal(t, httpresp.StatusCode, 202)

	event := eventResponse.GetEvent()
	status := eventResponse.GetStatus()
	assert.Equal(t, status, "ok")
	assert.Equal(t, event.GetTitle(), testEvent.GetTitle())
	assert.Equal(t, event.GetText(), testEvent.GetText())
	assert.Assert(t, event.GetUrl() != "")

	const retries = 5
	retry := 0
	var fetchedEvent datadog.Event
	for retry < retries {
		retry = retry + 1
		// Check event existence
		fetchedEventResponse, httpresp, err := TESTAPICLIENT.EventsApi.GetEvent(TESTAUTH, event.GetId()).Execute()
		if err != nil {
			if retry < retries {
				time.Sleep(5000 * time.Millisecond)
				t.Logf("Retry %d", retry)
				continue
			}
			t.Errorf("Error fetching Event %v: Response %v: %v", event.GetId(), err.(datadog.GenericOpenAPIError).Body(), err)
		}
		fetchedEvent = fetchedEventResponse.GetEvent()
		assert.Equal(t, httpresp.StatusCode, 200)
		assert.Equal(t, event.GetTitle(), fetchedEvent.GetTitle())
		assert.Equal(t, event.GetText(), fetchedEvent.GetText())
		// not the same!!! assert.Equal(t, event.GetUrl(), fetchedEvent.GetUrl())
		assert.Assert(t, fetchedEvent.GetUrl() != "")
		break
	}

	// Find our event in the full list
	start := event.GetDateHappened() - 10
	end := start + 20
	// eventListResponse, httpresp, err := TESTAPICLIENT.EventsApi.ListEvents(TESTAUTH).Start(start).End(end).Priority("normal").Sources("datadog-api-client-go").Tags("test,client:go").Unaggregated(true).Execute()
	eventListResponse, httpresp, err := TESTAPICLIENT.EventsApi.ListEvents(TESTAUTH).Start(start).End(end).Priority("normal").Sources("datadog-api-client-go").Tags("test,client:go").Unaggregated(true).Execute()
	if err != nil {
		t.Errorf("Error fetching events: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, httpresp.StatusCode, 200)

	events := eventListResponse.GetEvents()
	assert.Assert(t, is.Contains(events, fetchedEvent))
}
