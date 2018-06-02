package main

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

const (
	rows = 101
	cols = 51
)

type MCAgent struct {
	lastAction []int
	numActions int
	numStates  int
	gamma      float64
	p          float64
	states     []int
	rewards    []float64
	episode    []stateActionPair
	V          *mat.VecDense
	num        *mat.VecDense
}

func (agent *MCAgent) Init() {
	agent.states = make([]int, 0)
	agent.rewards = make([]float64, 0)
	agent.V = mat.NewVecDense(agent.numStates+1, nil)
	agent.num = mat.NewVecDense(agent.numStates+1, nil)
	for i := 0; i < agent.numStates+1; i++ {
		agent.num.SetVec(i, 1)
	}
}

func (agent *MCAgent) Start(thisObservation []int) []int {
	action := agent.chooseAction()
	agent.states = make([]int, 1)
	agent.states[0] = thisObservation[0]
	return action
}

func (agent *MCAgent) Step(reward float64, thisObservation []int) []int {
	s := thisObservation[0]
	action := agent.chooseAction()
	agent.states = append(agent.states, s)
	agent.rewards = append(agent.rewards, reward)
	return action
}

func (agent *MCAgent) End(reward float64) {
	agent.rewards = append(agent.rewards, reward)
	returnSoFar := 0.0
	visitedStates := mat.NewVecDense(agent.numStates+1, nil)
	returns := mat.VecDenseCopyOf(agent.V.TVec())
	for cnt := len(agent.states) - 1; cnt >= 0; cnt-- {
		state := agent.states[cnt]
		visitedStates.SetVec(state, 1)
		returnSoFar *= agent.gamma
		returnSoFar += agent.rewards[cnt]
		returns.SetVec(state, returnSoFar)
	}
	td := mat.NewVecDense(agent.numStates+1, nil)
	agent.num.AddVec(agent.num, visitedStates)
	returns.SubVec(returns, agent.V.TVec())

	td.DivElemVec(returns.TVec(), agent.num.TVec())
	agent.V.AddVec(agent.V, td.TVec())
}

func (agent *MCAgent) Cleanup() {
}

// TODO
func (agent *MCAgent) Message(inMessage string) interface{} {
	return nil
}

func NewMCAgent(n int, gamma float64, p float64) *MCAgent {
	return &MCAgent{numStates: n, gamma: gamma}
}

func (agent *MCAgent) chooseAction() []int {
	var action []int
	toss := rand.Float64()
	if toss < agent.p {
		action = []int{0}
	} else {
		action = []int{1}
	}
	return action
}
