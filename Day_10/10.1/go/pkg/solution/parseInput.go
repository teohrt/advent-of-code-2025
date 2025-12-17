package solution

import (
	"log"
	"solution/pkg/lineIterator"
	"strconv"
	"strings"
)

type Machine struct {
	lights  []string
	buttons [][]int
	joltage []int
}

func parseInput(filePath string) []Machine {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	var machines []Machine
	for iterator.Next() {
		line := iterator.Line()
		machine := parseMachine(line)
		machines = append(machines, machine)
	}
	return machines
}

func parseMachine(line string) Machine {
	machine := Machine{}

	// Parse lights pattern [.##.]
	lightsStart := strings.Index(line, "[")
	lightsEnd := strings.Index(line, "]")
	if lightsStart != -1 && lightsEnd != -1 {
		lightsStr := line[lightsStart+1 : lightsEnd]
		machine.lights = strings.Split(lightsStr, "")
	}

	// Parse buttons (3) (1,3) (2)...
	// Find all content between parentheses
	buttonStart := lightsEnd + 1
	joltageStart := strings.Index(line, "{")
	if joltageStart == -1 {
		joltageStart = len(line)
	}

	buttonSection := line[buttonStart:joltageStart]
	machine.buttons = parseButtons(buttonSection)

	// Parse joltage {3,5,4,7}
	if joltageStart != len(line) {
		joltageEnd := strings.Index(line, "}")
		if joltageEnd != -1 {
			joltageStr := line[joltageStart+1 : joltageEnd]
			machine.joltage = parseNumbers(joltageStr)
		}
	}

	return machine
}

func parseButtons(section string) [][]int {
	var buttons [][]int

	// Find all groups in parentheses
	inParen := false
	current := ""

	for _, ch := range section {
		if ch == '(' {
			inParen = true
			current = ""
		} else if ch == ')' {
			if inParen && current != "" {
				buttonIndices := parseNumbers(current)
				buttons = append(buttons, buttonIndices)
			}
			inParen = false
			current = ""
		} else if inParen {
			current += string(ch)
		}
	}

	return buttons
}

func parseNumbers(s string) []int {
	var numbers []int
	parts := strings.Split(s, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			num, err := strconv.Atoi(part)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	return numbers
}
