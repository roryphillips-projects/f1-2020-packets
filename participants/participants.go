package participants

import "github.com/roryphillips-projects/f1-2020-packets/common"

// Packet participants data
// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name. If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.
// N.B. on Xbox One, the names will always be the driver name, on PS4 the name will be the LAN name if playing a LAN game, otherwise it will be the driver name.
// The array should be indexed by vehicle index.
// Frequency: Every 5 seconds
// Size: 1213 bytes
type Packet struct {
	// Header packet header
	Header common.Header `json:"header"`
	// Number of active cars in the data - should match number of cars on HUD
	NumCars uint8 `json:"num_cars" packet:"0"`
	// Participant data
	Participants [22]ParticipantData `json:"participants" packets:"0"`

}

type ParticipantData struct {
	AIControlled bool
	DriverID uint8
	TeamID uint8
	RaceNumber uint8
	Nationality uint8
	Name [48]rune
}

