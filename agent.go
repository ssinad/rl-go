package main

import(
	"math/rand"
)

type SimpleAgent struct{
	last_action []int
	num_actions int
	Value_function []float64
}

func (agent *SimpleAgent) Init(){
	agent.num_actions = 10
	agent.Value_function = []float64{0.1, 0.2, 0.3}
	agent.last_action = []int{}
}

func (agent *SimpleAgent) Start(this_observation []int) []int{
	agent.last_action[0] = rand.Intn(agent.num_actions)
	local_action := []int{0}
	local_action[0] = rand.Intn(agent.num_actions)
	return local_action
}

func (agent *SimpleAgent) Step(reward float64, this_observation []int) []int{
	local_action := []int{0}
	local_action[0] = rand.Intn(agent.num_actions)
	agent.last_action = local_action
	return agent.last_action
}

func (agent *SimpleAgent) End(reward float64){
}

func (agent *SimpleAgent) Cleanup(){
}

// TODO
func (agent *SimpleAgent) Message(inMessage string) interface{}{
	return nil
}

func NewSimpleAgent() *SimpleAgent{
	return &SimpleAgent{}
}