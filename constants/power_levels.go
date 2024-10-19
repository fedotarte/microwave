package constants

type PowerLevel int

const (
	PowerLevel600  PowerLevel = 600
	PowerLevel800  PowerLevel = 800
	PowerLevel1200 PowerLevel = 1200
)

var ValidPowerLevels = [3]PowerLevel{
	PowerLevel600,
	PowerLevel800,
	PowerLevel1200,
}
