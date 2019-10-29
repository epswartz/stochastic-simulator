package main

import (
	"fmt"

	"github.com/epswartz/stochastic-simulator/pkg/model"
)

func main() {
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
		StartState: startState,
	}
}
