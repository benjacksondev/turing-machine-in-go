# turing-machine-in-go

A Turing Machine in Go to complement further learning in *CETM70 - Computer Science Principles And Practice* class at SU (Sunderland University).

The goal is to internalise some of the concepts surrounding a Turing Machine (TM)

**Features** 
✅ Is Even Program (Program to determine if given input is even)
✅ Accepts tape as stdin
✅ Prints tape to stdout

## Getting Started

To run the project, with Go installed, simply run `go run .`.

> Tape must end with an empty symbol for the Turing Machine to detect when it should stop processing



## Explanation

The Turing Machine comprises Tape, Head, State and Transitions. The tape is an array of Symbols that can be read and written by the Turing Machine. The Turing Machine can move along the tape one place to the right or the left while storing its position on the tape in the head. The Turing Machine's potential operations can be represented wholly in quintuplets. For Example, the quintuplet Q 0: □ → E states when State is Q and the value on the tape is 0, overwrite 0 with □ (blank) move the head to the right (□) and transition to State E. The quintuplet Q 0: □ * E similarly tells the Turing Machine to overwrite 0 with □ and allow the head to remain in its current position and finally transition to state E. This particular Turing Machine is configured to determine whether or not a number (in binary) is even.
