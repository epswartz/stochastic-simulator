package model

import "fmt"

// Represents the entire model
type Model struct {
	StartState State // Start state - the model begins here
}

// A single simulated world state for a model.
type modelState struct {
	currentState *State
}

// A single state of the automaton
// If the transitions slice is nil or empty, it's an end state used by RunToEnd
type State struct {
	Name        string
	Transitions []Transition // Target states and corresponding probabilities
}

// An edge from one state to another
type Transition struct {
	Probability float64 // The probability that this transition is taken
	Destination *State  // State that this transition goes to
}

// Checks whether a state is valid
// - all probabilities must sum to one
// - all destinations must be unique
// If it's an end state, there are no transitions
// Returns whether it's valid, and the reason, if it isn't.
func validateState(s State) (bool, string) {
	fmt.Println("UNIMPLEMENTED FUNCTION")
	return true, ""
}

// Runs a model to its end
// reps - number of times to simulate the model
// Returns a map of state names the model finished in, and the number of times the model finished in that state.
// If any of the singular runs fail, the entire thing returns an error.
func (m Model) RunToEnd(reps int) (map[string]int, err) {
	// TODO call RunToEndSingle
}

func (m Model) RunToEndSingle() (State, err) {
	// TODO
}
