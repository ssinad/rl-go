package main

type ToyEnvironment struct {
	currentState []int
	numStates    int
	isTerminal   bool
}

func (env *ToyEnvironment) Init() {
	env.currentState = make([]int, 1)
}

func (env *ToyEnvironment) Start() []int {
	env.currentState = make([]int, 1)
	env.currentState[0] = 0
	return env.currentState
}

func (env *ToyEnvironment) Step(thisAction []int) (float64, []int, bool) {
	env.currentState[0] += thisAction[0]
	reward := 0.0
	env.isTerminal = false
	if env.currentState[0] == env.numStates {
		env.isTerminal = true
		env.currentState = nil
		reward = 1.0
	}
	return reward, env.currentState, env.isTerminal
}

func (env *ToyEnvironment) Cleanup() {
}

// TODO
func (env *ToyEnvironment) Message(inMessage string) interface{} {
	return nil
}

func NewToyEnvironment(numStates int) *ToyEnvironment {
	return &ToyEnvironment{numStates: numStates}
}
