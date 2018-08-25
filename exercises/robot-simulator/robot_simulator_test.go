// +build step1 !step2,!step3

package robot

// Tests are separated into 3 steps.
//
// Run all tests with `go test` or run specific tests with the -tags option.
// Examples,
//
//    go test                      # run all tests
//    go test -tags step1          # run just step 1 tests.
//    go test -tags 'step1 step2'  # run step1 and step2 tests
//
// This source file contains step 1 tests only.  For other tests see
// robot_simulator_step2_test.go and robot_simulator_step3_test.go.
//
// You are given the source file defs.go which defines a number of things
// the test program requires.  It is organized into three sections by step.
//
// To complete step 1 you will define Right, Left, Advance, N, S, E, W,
// and Dir.String.  Complete step 1 before moving on to step 2.

import (
	"fmt"
	"testing"
)

var oneRobotTestCases = []struct {
	name string

	initX, initY int
	initFacing   Direction
	cmds         string

	finalX, finalY int
	finalFacing    Direction
}{
	{
		name:        "turn right from facing North",
		initX:       0,
		initY:       0,
		initFacing:  North,
		cmds:        "R",
		finalX:      0,
		finalY:      0,
		finalFacing: East,
	},
	{
		name:        "turn right from facing East",
		initX:       0,
		initY:       0,
		initFacing:  East,
		cmds:        "R",
		finalX:      0,
		finalY:      0,
		finalFacing: South,
	},
	{
		name:        "turn right from facing South",
		initX:       0,
		initY:       0,
		initFacing:  South,
		cmds:        "R",
		finalX:      0,
		finalY:      0,
		finalFacing: West,
	},
	{
		name:        "turn right from facing West",
		initX:       0,
		initY:       0,
		initFacing:  West,
		cmds:        "R",
		finalX:      0,
		finalY:      0,
		finalFacing: North,
	},
	{
		name:        "turn left from facing North",
		initX:       0,
		initY:       0,
		initFacing:  North,
		cmds:        "L",
		finalX:      0,
		finalY:      0,
		finalFacing: West,
	},
	{
		name:        "turn left from facing East",
		initX:       0,
		initY:       0,
		initFacing:  East,
		cmds:        "L",
		finalX:      0,
		finalY:      0,
		finalFacing: North,
	},
	{
		name:        "turn left from facing South",
		initX:       0,
		initY:       0,
		initFacing:  South,
		cmds:        "L",
		finalX:      0,
		finalY:      0,
		finalFacing: East,
	},
	{
		name:        "turn left from facing West",
		initX:       0,
		initY:       0,
		initFacing:  West,
		cmds:        "L",
		finalX:      0,
		finalY:      0,
		finalFacing: South,
	},
	{
		name:        "advance from facing North",
		initX:       0,
		initY:       0,
		initFacing:  North,
		cmds:        "A",
		finalX:      0,
		finalY:      1,
		finalFacing: North,
	},
	{
		name:        "advance from facing East",
		initX:       0,
		initY:       0,
		initFacing:  East,
		cmds:        "A",
		finalX:      1,
		finalY:      0,
		finalFacing: East,
	},
	{
		name:        "advance from facing South",
		initX:       0,
		initY:       0,
		initFacing:  South,
		cmds:        "A",
		finalX:      0,
		finalY:      -1,
		finalFacing: South,
	},
	{
		name:        "advance from facing West",
		initX:       0,
		initY:       0,
		initFacing:  West,
		cmds:        "A",
		finalX:      -1,
		finalY:      0,
		finalFacing: West,
	},
	{
		name:        "multple instructions 1",
		initX:       0,
		initY:       0,
		initFacing:  North,
		cmds:        "LAAARALA",
		finalX:      -4,
		finalY:      1,
		finalFacing: West,
	},
	{
		name:        "multple instructions 2",
		initX:       2,
		initY:       -7,
		initFacing:  East,
		cmds:        "RRAAAAALA",
		finalX:      -3,
		finalY:      -8,
		finalFacing: South,
	},
	{
		name:        "multple instructions 3",
		initX:       8,
		initY:       4,
		initFacing:  South,
		cmds:        "LAAARRRALLLL",
		finalX:      11,
		finalY:      5,
		finalFacing: North,
	},
}

func TestOneRobot(t *testing.T) {
	for _, tc := range oneRobotTestCases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRobot(tc.initX, tc.initY, tc.initFacing)

			for _, cmd := range tc.cmds {
				switch Command(cmd) {
				default:
					panic(fmt.Sprintf("unknown command: %v", cmd))
				case TurnLeft:
					r.TurnLeft()
				case TurnRight:
					r.TurnRight()
				case Advance:
					r.Advance()
				}
			}

			if r.X != tc.finalX || r.Y != tc.finalY {
				t.Errorf("different final position; want= {x= %d, y= %d}; got= {x= %d, y= %d}", tc.finalX, tc.finalY, r.X, r.Y)
			}
			if r.Facing != tc.finalFacing {
				t.Errorf("different final facing; want= %v; got= %v", tc.finalFacing, r.Facing)
			}
		})
	}
}

func TestNewRobot(t *testing.T) {
	actual := NewRobot(10, -10, West)
	expected := Robot{
		X:      10,
		Y:      -10,
		Facing: West,
	}

	if want, got := *actual, expected; want != got {
		t.Errorf("different result; want= %#v; got= %#v", want, got)
	}
}
