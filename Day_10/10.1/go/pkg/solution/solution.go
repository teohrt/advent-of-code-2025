package solution

import (
	"math"
)

/*
High level overview:
--------------------------------
parse input machines

result_total = 0

for every machine:
	buttonMap = create button map
	iterative DFS over button map
		update result_total

return result_total


Deeper dive:
--------------------------------

Example machine:
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}

map:
0: (0, 2), (0, 1)
1: (1, 3), (0, 1)
2: (2), (2, 3), (0, 2)
3: (3), (1,3), (2, 3)

iterative DFS:

state = [0,0,0,0]
desires = [1, 2]
pressCount = 0

min_count = infinity
stack = [(state, desires, pressCount)]
while stack
	state, desires, pressCount = stack.pop()
	if pressCount > min_count:
		continue
	if desires are satisfied:
		update min_count
		continue
	else:
		for every desire:
			for every associated button:
				push new state, desires & pressCount onto the stack

return min_count
*/

func Solve(filePath string) int {
	machines := parseInput(filePath)

	result := 0
	for _, machine := range machines {
		result += getMinPresses(machine)
	}

	return result
}

type Item struct {
	lightStates []bool
	pressCount  int
}

func getMinPresses(machine Machine) int {
	buttonMap := make(map[int][][]int)
	for _, b := range machine.buttons {
		for _, button := range b {
			buttonMap[button] = append(buttonMap[button], b)
		}
	}

	minPresses := int(float64(math.Inf(1)))
	stack := []Item{
		{
			lightStates: make([]bool, len(machine.desiredLights)), // initially all lights are off
			pressCount:  0,
		},
	}
	for len(stack) > 0 {
		var popped Item
		popped, stack = stack[0], stack[1:]

		// if the number of presses is greater than the minimum, skip further processing and continue
		// this is a pruning optimization - if we've already found a better solution, we don't need to continue
		if popped.pressCount > minPresses {
			continue
		}

		// determine the lights that need to be flipped to achieve the desired state
		// if none, we've found the desired state, can update the minimum and continue
		lightsToFlip := []int{}
		for idx, desiredLightValue := range machine.desiredLights {
			if desiredLightValue != popped.lightStates[idx] {
				lightsToFlip = append(lightsToFlip, idx)
			}
		}
		foundDesiredState := len(lightsToFlip) == 0
		if foundDesiredState {
			minPresses = popped.pressCount
			continue
		}

		// push new states onto the stack
		for _, idxToFlip := range lightsToFlip {
			for _, button := range buttonMap[idxToFlip] {
				// "Press the button" & update the light states associated
				newLightStates := make([]bool, len(popped.lightStates))
				copy(newLightStates, popped.lightStates)
				for _, lightIdx := range button {
					newLightStates[lightIdx] = !newLightStates[lightIdx]
				}
				item := Item{
					lightStates: newLightStates,
					pressCount:  popped.pressCount + 1,
				}
				stack = append(stack, item)
			}
		}
	}

	return int(minPresses)
}
