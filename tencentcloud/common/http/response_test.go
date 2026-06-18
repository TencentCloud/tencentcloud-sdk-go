package common

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// buildSSEResponse wraps body as a text/event-stream http.Response with a
// cancellable request context (parseFromSSE reads hr.Request.Context()).
func buildSSEResponse(body string) *http.Response {
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://x", nil)
	return &http.Response{
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"text/event-stream"}},
		Body:       ioutil.NopCloser(strings.NewReader(body)),
		Request:    req,
	}
}

// drain collects every event from r.Events until channel close.
func drain(r *BaseSSEResponse) []SSEvent {
	var out []SSEvent
	for ev := range r.Events {
		out = append(out, ev)
	}
	return out
}

// TestParseFromSSE_LargeLine verifies a single data line larger than the
// stdlib bufio.Scanner default (64KB) parses successfully — regression guard
// for the silent ErrTooLong bug fixed by raising the buffer cap.
func TestParseFromSSE_LargeLine(t *testing.T) {
	const sz = 32 * 1024 // 32KB, below 64KB default cap
	payload := strings.Repeat("x", sz)
	body := "data: " + payload + "\n\n"

	resp := &BaseSSEResponse{}
	if err := parseFromSSE(buildSSEResponse(body), resp); err != nil {
		t.Fatalf("parseFromSSE returned error: %v", err)
	}
	events := drain(resp)
	if len(events) != 1 {
		t.Fatalf("want 1 event, got %d", len(events))
	}
	ev := events[0]
	if ev.Err != nil {
		t.Fatalf("unexpected event error: %v", ev.Err)
	}
	if len(ev.Data) != sz {
		t.Fatalf("want data len %d, got %d", sz, len(ev.Data))
	}
	if !bytes.Equal(ev.Data, []byte(payload)) {
		t.Fatal("data payload mismatch")
	}
}

// TestParseFromSSE_ExceedsCap shrinks SSEScannerBufferMaxBytes so a moderate
// line overflows. Asserts the scanner error surfaces via the channel instead
// of silent truncation.
func TestParseFromSSE_ExceedsCap(t *testing.T) {
	prev := SSEScannerBufferMaxBytes
	SSEScannerBufferMaxBytes = 4 * 1024
	defer func() { SSEScannerBufferMaxBytes = prev }()

	body := "data: " + strings.Repeat("x", 8*1024) + "\n\n"

	resp := &BaseSSEResponse{}
	if err := parseFromSSE(buildSSEResponse(body), resp); err != nil {
		t.Fatalf("parseFromSSE returned error: %v", err)
	}
	events := drain(resp)

	var sawErr bool
	for _, ev := range events {
		if ev.Err != nil {
			sawErr = true
			break
		}
	}
	if !sawErr {
		t.Fatalf("expected scanner.Err to surface via channel, got events=%+v", events)
	}
}

// TestParseFromSSE_ConfigurableCap raises the cap and confirms a line that
// would overflow the default-shrunk cap still parses.
func TestParseFromSSE_ConfigurableCap(t *testing.T) {
	prev := SSEScannerBufferMaxBytes
	SSEScannerBufferMaxBytes = 16 * 1024 * 1024
	defer func() { SSEScannerBufferMaxBytes = prev }()

	const sz = 4 * 1024 * 1024
	payload := strings.Repeat("y", sz)
	body := "data: " + payload + "\n\n"

	resp := &BaseSSEResponse{}
	if err := parseFromSSE(buildSSEResponse(body), resp); err != nil {
		t.Fatalf("parseFromSSE returned error: %v", err)
	}
	events := drain(resp)
	if len(events) != 1 || events[0].Err != nil {
		t.Fatalf("want 1 clean event, got %+v", events)
	}
	if len(events[0].Data) != sz {
		t.Fatalf("want data len %d, got %d", sz, len(events[0].Data))
	}
}

// TestParseFromSSE_MultiField exercises event/data/id/retry parsing — keeps
// the buffer-cap regression tests honest by proving the parser still handles
// normal SSE framing alongside the large-line cases.
func TestParseFromSSE_MultiField(t *testing.T) {
	body := "event: msg\ndata: hello\ndata: world\nid: 42\nretry: 1500\n\n"

	resp := &BaseSSEResponse{}
	if err := parseFromSSE(buildSSEResponse(body), resp); err != nil {
		t.Fatalf("parseFromSSE returned error: %v", err)
	}
	events := drain(resp)
	if len(events) != 1 {
		t.Fatalf("want 1 event, got %d", len(events))
	}
	ev := events[0]
	if ev.Event != "msg" || string(ev.Data) != "hello\nworld" || ev.Id != "42" || ev.Retry != 1500 {
		t.Fatalf("event mismatch: %+v", ev)
	}
}
