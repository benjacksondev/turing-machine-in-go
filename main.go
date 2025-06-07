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
			// This is a simple program to deterimine if provided binary input is even or odd
			//
			// The start state Q - probably not required but doesn't make
			// sense to start at state E (Even) or O (Odd) when it is not
			// yet determined.
			//
			// if the value on the tape is even move to state 1 (for Even)
			// else the value on the tape is odd move to state O (for Odd)
			// We don't need to write anything to the tape yet
			"Q": {
				"0": {NextState: "E", Write: " ", Move: 1}, // Q 0: □ → E (-> states move head to the right)
				"1": {NextState: "O", Write: " ", Move: 1}, // Q 1: □ → O
			},
			// The even state E
			// if the value on the tape is even move to state 1 (for Even)
			// elseif the value on the tape is odd move to state O (for Odd)
			// else the value on the tape is blank write 1 (is even)
			"E": {
				"0": {NextState: "E", Write: " ", Move: 1}, // E 0: □ → E
				"1": {NextState: "O", Write: " ", Move: 1}, // E 1: □ → O
				" ": {NextState: "F", Write: "1", Move: 0}, // E □: 1 * F (* states do not move the head)
			},
			// The odd state O
			// if the value on the tape is even move to state 1 (for Even)
			// elseif the value on the tape is odd move to state O (for Odd)
			// else the value on the tape is blank write 0 (is odd)
			"O": {
				"0": {NextState: "E", Write: " ", Move: 1}, // O 0: □ → E
				"1": {NextState: "O", Write: " ", Move: 1}, // O 1: □ → O
				" ": {NextState: "F", Write: "0", Move: 0}, // O □: 0 * F
			},
		},
	}

	tm.Run()
}
