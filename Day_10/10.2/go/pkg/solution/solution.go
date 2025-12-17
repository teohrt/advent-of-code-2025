package solution

import (
	"fmt"
)

/*
High level overview:
--------------------------------
parse input machines

result_total = 0

for every machine:
	buttonMap = create button map
	iterative BFS over button states
		update result_total

return result_total
*/

func Solve(filePath string) int {
	machines := parseInput(filePath)

	result := 0
	for i, machine := range machines {
		minPresses := getMinPresses(machine)
		result += minPresses
		fmt.Printf("machine %d, min presses: %d\n", i, minPresses)
	}

	return result
}

type Item struct {
	presses       int
	counterStates []int
}

// BFS to find the minimum number of presses to achieve the desired state
func getMinPresses(machine Machine) int {
	targetState := machine.joltage
	numCounters := len(targetState)

	// Helper to convert []int to string for map key
	stateToString := func(states []int) string {
		b := make([]byte, numCounters)
		for i, s := range states {
			b[i] = byte(s)
		}
		return string(b)
	}

	initialStates := make([]int, numCounters)
	initialStateStr := stateToString(initialStates)
	targetStateStr := stateToString(targetState)

	if initialStateStr == targetStateStr {
		return 0
	}

	queue := []Item{
		{
			counterStates: initialStates,
			presses:       0,
		},
	}

	visited := make(map[string]bool)
	visited[initialStateStr] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, button := range machine.buttons {
			// Create new state by toggling lights listed in the button
			nextStates := make([]int, numCounters)
			copy(nextStates, current.counterStates)
			for _, counterIdx := range button {
				nextStates[counterIdx] = nextStates[counterIdx] + 1
			}

			// If any counter state is greater than the target state, we can skip
			skip := false
			for counterIdx, counterState := range nextStates {
				if counterState > targetState[counterIdx] {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			nextStateStr := stateToString(nextStates)
			if nextStateStr == targetStateStr {
				return current.presses + 1
			}

			// If we've already visited the next state, it's guarenteed that
			// we got there with <= presses than whatever we have now, so we can skip
			if !visited[nextStateStr] {
				visited[nextStateStr] = true
				queue = append(queue, Item{
					counterStates: nextStates,
					presses:       current.presses + 1,
				})
			}
		}
	}

	return 0 // Should not happen based on problem description
}
