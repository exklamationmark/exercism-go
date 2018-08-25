package robot

import "fmt"

// NewRobot creates a new robot with an initial position & facing direction.
func NewRobot(x, y int, direction Direction) *Robot {
	return &Robot{
		X:      x,
		Y:      y,
		Facing: direction,
	}
}

// Advance updates the robot's coordinate depending on where it is facing.
func (r *Robot) Advance() {
	switch r.Facing {
	default:
		panic(fmt.Sprintf("unknown direction: %v", r.Facing))
	case North:
		r.Y = r.Y + 1
	case South:
		r.Y = r.Y - 1
	case East:
		r.X = r.X + 1
	case West:
		r.X = r.X - 1
	}

	// TODO(): handle integer overflow
}

// TurnLeft changes the direction the robot is facing by 90 degree counter-clockwise.
func (r *Robot) TurnLeft() {
	newDirection := (r.Facing - 1 + noDirection) % noDirection
	r.Facing = newDirection
}

// TurnRight changes the direction the robot is facing by 90 degree clockwise.
func (r *Robot) TurnRight() {
	newDirection := (r.Facing + 1) % noDirection
	r.Facing = newDirection
}
