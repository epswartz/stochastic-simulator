package model

import (
	"errors"
	"fmt"
	"math/rand"

	u "github.com/epswartz/stochastic-simulator/pkg/util"
)

/*
TODO
	0. Need a list of states in Model type so they can all be validated before it's run
	1. User-defined step limit and time limit for runs
	2. Keep history of states
*/
// Represents the entire model
type Model struct {
	StartState *State // Start state - the model begins here
}

// A single simulated world state for a model.
type modelState struct {
	currentState *State
}

// Moves a modelState to the next state, and also returns that state for convenience.
// Relies on the state having already been validated with validateState.
func (ms modelState) nextState() *State {
	r := rand.Float64()
	for _, t := range ms.currentState.Transitions {
		r -= t.Probability
		if r < 0 {
			ms.currentState = t.Destination
			return t.Destination
		}
	}

	// Returns the last one if we never get below zero, as a way to deal with float inaccuracy
	ms.currentState = ms.currentState.Transitions[len(ms.currentState.Transitions)-1].Destination
	return ms.currentState.Transitions[len(ms.currentState.Transitions)-1].Destination

}

// Used to pass result of single simulation through channel.
type simResponse struct {
	endState *State
	err      error
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
// Returns error containing a reason, if it isn't valid.
func validateState(s *State) error {
	if len(s.Transitions) == 0 {
		return errors.New("Empty transitions slice")
	}

	destinations := map[*State]struct{}{}

	var probabilitySum float64 = 0

	for _, t := range s.Transitions {
		if _, ok := destinations[t.Destination]; ok {
			return errors.New("Non-unique destination")
		}
		destinations[t.Destination] = struct{}{}
		probabilitySum += t.Probability
	}

	if !u.FloatEqual(probabilitySum, 0.0) {
		return errors.New("Probability values do not sum to one")
	}
	fmt.Println("UNIMPLEMENTED FUNCTION")
	return nil
}

// Runs a model to its end
// reps - number of times to simulate the model
// maxSteps - maximum number of steps for each simulation to take before aborting
// Returns a map of state names the model finished in, and the number of times the model finished in that state.
// If any of the singular runs fail, the entire thing returns an error.
func (m Model) RunToEnd(reps, maxSteps int) (map[string]int, error) {
	// TODO validate all the states.
	// TODO call RunToEndSingle
	return nil, nil
}

// TODO comment
func (m Model) RunToEndSingle(maxSteps int, resChan chan simResponse) {
	// TODO Start in start state, randomly take transitions until you get an end state.
	ms := modelState{
		currentState: m.StartState,
	}

	for i := 0; i < maxSteps; i++ {
		err := validateState(ms.currentState)
		if err != nil {
			resChan <- simResponse{
				endState: ms.currentState,
				err:      err,
			}
			return
		}
		ms.nextState()
	}

}
