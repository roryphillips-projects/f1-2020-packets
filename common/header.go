package common

// Header format for F1 2020 telemetry data
type Header struct {
	// Header specifying packet format
	// i.e. 2020
	PacketFormat uint16 `json:"packet_format" packet:"0"`
	// GameMajorVersion major version identifier I.E. "X.00"
	GameMajorVersion uint8 `json:"game_major_version" packet:"1"`
	// GameMinorVersion minor version identifier "1.XX"
	GameMinorVersion uint8 `json:"game_minor_version" packet:"2"`
	// PacketVersion version of the packet type, all start from 1
	PacketVersion uint8 `json:"packet_version" packet:"3"`
	// PacketID type of the packet returned
	PacketID PacketID `json:"packet_id" packet:"4"`
	// SessionUID unique identifier for the session
	SessionUID uint64 `json:"session_uid" packet:"5"`
	// SessionTime timestamp of the session
	SessionTime float32 `json:"session_time" packet:"6"`
	// FrameIdentifier the frame the data was retrieved on
	FrameIdentifier uint32 `json:"frame_identifier" packet:"7"`
	// PlayerCarIndex the player's car in the array
	PlayerCarIndex uint8 `json:"player_car_index" packet:"8"`
	// SecondaryPlayerCarIndex the secondary player's car in the array (splitscreen)
	// 255 if no second player
	SecondaryPlayerCarIndex uint8 `json:"secondary_player_car_index" packet:"9"`
}

