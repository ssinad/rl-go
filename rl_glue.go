package main

type Agent interface {
	Init()
	Start([]int) []int
	Step(float64, []int) []int
	End(float64)
	Cleanup()
	Message(string) interface{}
}

type Environment interface{
	Init()
	Start() []int
	Step([]int) (float64, []int, bool)
	Cleanup()
	Message(string) interface{}
}

type RLGlue struct{
	env Environment
	agent Agent
	total_reward float64
	num_steps int
	num_episodes int
	last_state []int
	last_action []int
}

func (rl *RLGlue) Init(){
	rl.env.Init()
	rl.agent.Init()
	rl.total_reward = 0.0
	rl.num_steps = 0
	rl.num_episodes = 0
}

func (rl *RLGlue) Start() ([]int, []int){
	rl.total_reward = 0.0
	rl.num_steps = 1
	rl.last_state = rl.env.Start()
	rl.last_action = rl.agent.Start(rl.last_state)
	return rl.last_state, rl.last_action
}

func (rl *RLGlue) Step() (float64, []int, []int, bool){
	this_reward, last_state, terminal := rl.env.Step(rl.last_action)
	rl.total_reward += this_reward

	if terminal{
		rl.num_episodes += 1
		rl.agent.End(this_reward)
		return this_reward, last_state, rl.last_action, terminal
	}else{
		rl.num_steps += 1
		rl.last_action = rl.agent.Step(this_reward, last_state)
		return this_reward, last_state, rl.last_action, terminal
	}
}

func (rl *RLGlue) Cleanup(){
	rl.env.Cleanup()
	rl.agent.Cleanup()
}

// TODO
func (rl *RLGlue) Message(inMessage string) interface{}{
	return nil
}

func (rl *RLGlue) Episode(max_steps_this_episode int) bool{
	is_terminal := false
	rl.Start()
	for !is_terminal && (max_steps_this_episode == 0 || rl.num_steps < max_steps_this_episode){
		_, _, _, is_terminal = rl.Step()
	}
	return is_terminal
}