package robot

// This file contains definitions used in the simulation.

//go:generate stringer -type=Direction
type Direction int

const (
	North Direction = iota
	East
	South
	West

	noDirection = 4
)

type Robot struct {
	X, Y   int
	Facing Direction
}

type Command byte

//go:generate stringer -type=Command
const (
	Advance   Command = 'A'
	TurnLeft  Command = 'L'
	TurnRight Command = 'R'
)
