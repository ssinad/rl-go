package main

import (
	"fmt"
	"math/rand"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type GamblerEnvironment struct {
	headProbability float64
	currentState    []int
	numStates       int
	isTerminal      bool
}

func (env *GamblerEnvironment) Init() {
	env.currentState = []int{0}
	env.headProbability = 0.55
}

func (env *GamblerEnvironment) Start() []int {
	env.currentState = []int{rand.Intn(env.numStates-1) + 1}
	return env.currentState
}

func (env *GamblerEnvironment) Step(thisAction []int) (float64, []int, bool) {
	if thisAction[0] < 1 || thisAction[0] > min(env.currentState[0], env.numStates-env.currentState[0]) {
		panic(fmt.Sprintf("Invalid action taken!!\nAction: %d\n Current State: %d", thisAction, env.currentState))
	}
	uniformRand := rand.Float64()
	if uniformRand < env.headProbability {
		env.currentState[0] += thisAction[0]
	} else {
		env.currentState[0] -= thisAction[0]
	}
	reward := 0.0
	env.isTerminal = false
	if env.currentState[0] == env.numStates+1 {
		env.isTerminal = true
		env.currentState = nil
		reward = 1.0
	} else if env.currentState[0] == 0 {
		env.isTerminal = true
		env.currentState = nil
	}
	return reward, env.currentState, env.isTerminal
}

func (env *GamblerEnvironment) Cleanup() {
}

// TODO
func (env *GamblerEnvironment) Message(inMessage string) interface{} {
	return nil
}

func NewGamblerEnvironment(numStates int, headProbability float64) *GamblerEnvironment {
	return &GamblerEnvironment{numStates: numStates, headProbability: headProbability}
}
