# stochastic-simulator

A really basic stochastic model (at the time I write this, I mean basically a probablistic state machine - I'm not actually well-versed enough to know what to call that) simulator.

## Usage
```
    // Seed the random number gen
	rand.Seed(time.Now().UTC().UnixNano())

    // Create end states
	endState1 := model.State{
		Name: "End1",
	}
	endState2 := model.State{
		Name: "End2",
	}

    // Create probabilistic transition list for the start state
    // (end states have no transitions)
	tl := []model.Transition{
		model.Transition{
			Probability: 0.5,
			Destination: &endState1,
		},
		model.Transition{
			Probability: 0.5,
			Destination: &endState2,
		},
	}

    // Create start state
	startState := model.State{
		Name:        "Start",
		Transitions: tl,
	}

    // Create a model object, give it the start state
	m := model.Model{
		StartState: &startState,
	}

    // Run the model 10 times
    // Each run is limited to 5 steps before failure
	responses := m.RunToEnd(10, 5)
	for _, r := range responses {
		fmt.Println(r.EndState.Name, r.Err)
	}

    // If you wanted stats on where the model would end, you could then count the end states from responses
```
