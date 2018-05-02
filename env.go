package main

import (
	"math/rand"
)

type observation struct {
	reward     float64
	state      []int
	isTerminal bool
}

type SimpleEnvironment struct {
	thisRewardObservation observation
	numStates             int
}

func (env *SimpleEnvironment) Init() {
	localObservation := []int{}
	env.numStates = 10
	env.thisRewardObservation = observation{0, localObservation, false}
}

func (env *SimpleEnvironment) Start() []int {
	return env.thisRewardObservation.state
}

func (env *SimpleEnvironment) Step(this_action []int) (float64, []int, bool) {
	theReward := rand.NormFloat64()
	return theReward, env.thisRewardObservation.state, false
}

func (env *SimpleEnvironment) Cleanup() {
}

// TODO
func (env *SimpleEnvironment) Message(inMessage string) interface{} {
	return nil
}

func NewSimpleEnvironment() *SimpleEnvironment {
	return &SimpleEnvironment{}
}
