package main

import (
	"math/rand"
)

type SimpleAgent struct {
	lastAction    []int
	numActions    int
	ValueFunction []float64
}

func (agent *SimpleAgent) Init() {
	agent.numActions = 10
	agent.ValueFunction = []float64{0.1, 0.2, 0.3}
	agent.lastAction = []int{}
}

func (agent *SimpleAgent) Start(this_observation []int) []int {
	agent.lastAction[0] = rand.Intn(agent.numActions)
	localAction := []int{0}
	localAction[0] = rand.Intn(agent.numActions)
	return localAction
}

func (agent *SimpleAgent) Step(reward float64, this_observation []int) []int {
	localAction := []int{0}
	localAction[0] = rand.Intn(agent.numActions)
	agent.lastAction = localAction
	return agent.lastAction
}

func (agent *SimpleAgent) End(reward float64) {
}

func (agent *SimpleAgent) Cleanup() {
}

// TODO
func (agent *SimpleAgent) Message(inMessage string) interface{} {
	return nil
}

func NewSimpleAgent() *SimpleAgent {
	return &SimpleAgent{}
}
