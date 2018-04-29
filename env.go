package main

import(
	"math/rand"
)

type observation struct{
	reward 	float64
	state   []int
	is_terminal bool
}

type SimpleEnvironment struct{
	this_reward_observation observation
	num_states int
}

func (env *SimpleEnvironment) Init(){
	local_observation := []int{}
	env.num_states = 10
	env.this_reward_observation = observation{0, local_observation, false}
}

func (env *SimpleEnvironment) Start() []int{
	return env.this_reward_observation.state
}

func (env *SimpleEnvironment) Step(this_action []int) (float64, []int, bool){
	the_reward := rand.NormFloat64()
	return the_reward, env.this_reward_observation.state, false
}

func (env *SimpleEnvironment) Cleanup(){
}

// TODO
func (env *SimpleEnvironment) Message(inMessage string) interface{}{
	return nil
}