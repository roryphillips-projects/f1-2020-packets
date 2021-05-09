package common

// PacketID Identifier for the type of packet parsed
type PacketID uint8

const (
	// PacketIDMotion Motion data for player's car - only sent while the player is in control
	PacketIDMotion PacketID = 0
	// PacketIDSession Data about the session - track, time left
	PacketIDSession PacketID = 1
	// PacketIDLapData Data about all the lap times of cars in the session
	PacketIDLapData PacketID = 2
	// PacketIDEvent Various notable events that happen during a session
	PacketIDEvent PacketID = 3
	// PacketIDParticipants List of participants in the session, mostly relevant for multiplayer
	PacketIDParticipants PacketID = 4
	// PacketIDCarSetups Details of car setups in the race
	PacketIDCarSetups PacketID = 5
	// PacketIDCarTelemetry Telemetry data for all cars
	PacketIDCarTelemetry PacketID = 6
	// PacketIDCarStatus Status data for all cars such as damage
	PacketIDCarStatus PacketID = 7
	// PacketIDFinalClassification Final classification confirmation at the end of a race
	PacketIDFinalClassification PacketID = 8
	// PacketIDLobbyInfo Information about players in a multiplayer lobby
	PacketIDLobbyInfo PacketID = 9
)

// SessionType type of session
type SessionType uint8

const (
	// SessionTypeUnknown unknown session type
	SessionTypeUnknown SessionType = 0
	// SessionTypePractice1 practice one
	SessionTypePractice1 SessionType = 1
	// SessionTypePractice2 practice two
	SessionTypePractice2 SessionType = 2
	// SessionTypePractice3 practice three
	SessionTypePractice3 SessionType = 3
	// SessionTypeShortPractice short practice
	SessionTypeShortPractice SessionType = 4
	// SessionTypeQualifying1 qualifying one
	SessionTypeQualifying1 SessionType = 5
	// SessionTypeQualifying2 qualifying two
	SessionTypeQualifying2 SessionType = 6
	// SessionTypeQualifying3 qualifying three
	SessionTypeQualifying3 SessionType = 7
	// SessionTypeShortQualifying short qualifying
	SessionTypeShortQualifying SessionType = 8
	// SessionTypeOneShotQualifying one shot qualifying
	SessionTypeOneShotQualifying SessionType = 9
	// SessionTypeRace1 race one
	SessionTypeRace1 SessionType = 10
	// SessionTypeRace2 race two
	SessionTypeRace2 SessionType = 11
	// SessionTypeTimeTrial time trial
	SessionTypeTimeTrial SessionType = 12
)

// WeatherType type of weather
type WeatherType uint8

const (
	//WeatherTypeClear clear
	WeatherTypeClear WeatherType = 0
	//WeatherTypeLightCloud light cloud
	WeatherTypeLightCloud WeatherType = 1
	//WeatherTypeOvercast overcast
	WeatherTypeOvercast WeatherType = 2
	//WeatherTypeLightRain light rain
	WeatherTypeLightRain WeatherType = 3
	//WeatherTypeHeavyRain heavy rain
	WeatherTypeHeavyRain WeatherType = 4
	//WeatherTypeStorm storm
	WeatherTypeStorm WeatherType = 5
)

type TrackType int8

const (
	// TrackTypeUnknown Unknown track
	TrackTypeUnknown TrackType = -1
)

type FormulaType uint8

const (
	// FormulaTypeF1Modern F1 Modern
	FormulaTypeF1Modern FormulaType = 0
	// FormulaTypeF1Classic F1 Classic
	FormulaTypeF1Classic FormulaType = 1
	// FormulaTypeF2 F2
	FormulaTypeF2 FormulaType = 2
	// FormulaTypeF1Generic F1 Generic
	FormulaTypeF1Generic FormulaType = 3
)

// ZoneFlag flag flown
type ZoneFlag int8

const (
	// ZoneFlagUnknown unknown flag
	ZoneFlagUnknown ZoneFlag = -1
	// ZoneFlagNone no flag type
	ZoneFlagNone ZoneFlag = 0
	// ZoneFlagGreen green flag
	ZoneFlagGreen ZoneFlag = 1
	// ZoneFlagBlue blue flag
	ZoneFlagBlue ZoneFlag = 2
	// ZoneFlagYellow yellow flag
	ZoneFlagYellow ZoneFlag = 3
	// ZoneFlagRed red flag
	ZoneFlagRed ZoneFlag = 4
)

type PitStatus uint8

const (
	// PitStatusNone Not in the pits
	PitStatusNone PitStatus = 0
	// PitStatusPitting Entering the pit area
	PitStatusPitting PitStatus = 1
	// PitStatusInPitArea In the pits
	PitStatusInPitArea PitStatus = 2
)

type Sector uint8

const (
	// Sector1 First sector
	Sector1 Sector = 0
	// Sector2 Second sector
	Sector2 Sector = 1
	// Sector3 Third sector
	Sector3 Sector = 2
)

type DriverStatus uint8

const (
	// DriverStatusInGarage in the garage
	DriverStatusInGarage DriverStatus = 0
	// DriverStatusFlyingLap flying lap
	DriverStatusFlyingLap DriverStatus = 1
	// DriverStatusInLap in lap
	DriverStatusInLap DriverStatus = 2
	// DriverStatusOutLap Out Lap
	DriverStatusOutLap DriverStatus = 3
	// DriverStatusOnTrack On track
	DriverStatusOnTrack DriverStatus = 4
)

type ResultStatus uint8

const (
	//ResultStatusInvalid Invalid
	ResultStatusInvalid ResultStatus = 0
	//ResultStatusInactive Inactive
	ResultStatusInactive ResultStatus = 1
	//ResultStatusActive Active
	ResultStatusActive ResultStatus = 2
	//ResultStatusFinished Finished
	ResultStatusFinished ResultStatus = 3
	//ResultStatusDisqualified Disqualified
	ResultStatusDisqualified ResultStatus = 4
	//ResultStatusNotClassified Not Classified
	ResultStatusNotClassified ResultStatus = 5
	//ResultStatusRetired Retired
	ResultStatusRetired ResultStatus = 6
)
