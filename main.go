package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type State string
type Symbol string

type Transition struct {
	NextState State
	Write     Symbol
	Move      int // -1 for left, 1 for right
}

type TuringMachine struct {
	Tape        []Symbol
	Head        int
	State       State
	Transitions map[State]map[Symbol]Transition
}

func (tm *TuringMachine) Step() bool {
	symbol := tm.Tape[tm.Head]
	trans, exists := tm.Transitions[tm.State][symbol]
	if !exists {
		return false // Halt if no transition exists
	}

	tm.Tape[tm.Head] = trans.Write
	tm.Head += trans.Move
	if tm.Head < 0 {
		tm.Head = 0 // Prevent moving out of bounds
	}
	tm.State = trans.NextState
	time.Sleep(500 * time.Millisecond) // Add delay for mechanical effect
	return true
}

func (tm *TuringMachine) Run() {
	for tm.Step() {
		tm.PrintTape()
	}
}

func (tm *TuringMachine) PrintTape() {
	for i, s := range tm.Tape {
		if i == tm.Head {
			fmt.Printf("[%s]", s)
		} else {
			fmt.Printf(" %s ", s)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Enter tape: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	tape := make([]Symbol, len(input))
	for i, char := range input {
		tape[i] = Symbol(string(char))
	}

	tm := TuringMachine{
		Tape:  tape,
		Head:  0,
		State: "Q",
		Transitions: map[State]map[Symbol]Transition{
			"Q": {
				"0": {NextState: "E", Write: " ", Move: 1},
				"1": {NextState: "O", Write: " ", Move: 1},
			},
			"E": {
				"0": {NextState: "E", Write: " ", Move: 1},
				"1": {NextState: "O", Write: " ", Move: 1},
				" ": {NextState: "F", Write: "1", Move: 0},
			},
			"O": {
				"0": {NextState: "E", Write: " ", Move: 1},
				"1": {NextState: "O", Write: " ", Move: 1},
				" ": {NextState: "F", Write: "0", Move: 0},
			},
		},
	}

	tm.Run()
}
