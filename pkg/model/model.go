package model

import (
	"errors"
	"math/rand"

	u "github.com/epswartz/stochastic-simulator/pkg/util"
)

/*
TODO
	0. Need a list of states in Model type so they can all be validated before it's run
	1. User-defined step limit and time limit for runs
	2. Keep history of states, return through run methods
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
func (ms *modelState) nextState() *State {
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
type SimResponse struct {
	EndState *State
	Err      error
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

	destinations := map[*State]struct{}{}

	var probabilitySum float64 = 0

	// Make sure all the transitions add to 1
	if s.Transitions != nil {
		for _, t := range s.Transitions {
			if _, ok := destinations[t.Destination]; ok {
				return errors.New("Non-unique destination")
			}
			destinations[t.Destination] = struct{}{}
			probabilitySum += t.Probability
		}
		if !u.FloatEqual(probabilitySum, 1.0) {
			return errors.New("Probability values do not sum to one")
		}
	}

	return nil
}

// Runs a model to its end
// reps - number of times to simulate the model
// maxSteps - maximum number of steps for each simulation to take before aborting
// Returns a map of state names the model finished in, and the number of times the model finished in that state.
// Returns a list of errors, from each simulation.
func (m Model) RunToEnd(reps, maxSteps int) []SimResponse {
	// TODO call RunToEndSingle

	resChan := make(chan SimResponse)

	ret := make([]SimResponse, 0, reps)

	for i := 0; i < reps; i++ {
		go m.RunToEndSingle(maxSteps, resChan)
	}

	for i := 0; i < reps; i++ {
		response := <-resChan
		ret = append(ret, response)
	}

	return ret
}

// Runs the model through a single run., puts result in the channel given to it.
// Runs either to an end state, or to the maxSteps limit.:w
func (m Model) RunToEndSingle(maxSteps int, resChan chan SimResponse) {
	ms := modelState{
		currentState: m.StartState,
	}

	for i := 0; i < maxSteps; i++ {
		err := validateState(ms.currentState)
		if err != nil { // If current state is invalid, we're done
			resChan <- SimResponse{
				EndState: ms.currentState,
				Err:      err,
			}
			return
		}

		if ms.currentState.Transitions == nil || len(ms.currentState.Transitions) == 0 { // If it's an end state
			resChan <- SimResponse{
				EndState: ms.currentState,
				Err:      nil,
			}
			return
		}

		ms.nextState()
	}
	resChan <- SimResponse{
		EndState: ms.currentState,
		Err:      errors.New("Reached max number of steps."),
	}
}
