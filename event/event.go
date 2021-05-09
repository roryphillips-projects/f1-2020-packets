package event

import "github.com/roryphillips-projects/f1-2020-packets/common"

// Packet event data
// This packet gives details of events that happen during the course of a session.
// Frequency: When the event occurs
// Size: 35 bytes
type Packet struct {
	// Header packet header
	Header common.Header `json:"header"`
}
