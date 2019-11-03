package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/epswartz/stochastic-simulator/pkg/model"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Constructing basic model")
	endState1 := model.State{
		Name: "End1",
	}
	endState2 := model.State{
		Name: "End2",
	}

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

	startState := model.State{
		Name:        "Start",
		Transitions: tl,
	}

	m := model.Model{
		StartState: &startState,
	}

	responses := m.RunToEnd(10, 5)
	for _, r := range responses {
		fmt.Println(r.EndState.Name, r.Err)
	}

}
