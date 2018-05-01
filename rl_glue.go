package main

// import(
// 	"reflect"
// )

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
	is_terminal bool
}

func array_copy(arr []int) []int{
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	return cpy
}

// func array_copy(arr []reflect.Type){
// 	cpy := reflect.ArrayOf(len(arr), reflect.TypeOf(arr))
// 	reflect.Copy(cpy, arr)
// 	return cpy
// }

func (rl *RLGlue) Init(){
	rl.env.Init()
	rl.agent.Init()
	rl.total_reward = 0.0
	rl.num_steps = 0
	rl.num_episodes = 0
	rl.is_terminal = false
}

func (rl *RLGlue) Start() ([]int, []int){
	rl.total_reward = 0.0
	rl.is_terminal = false
	rl.num_steps = 1
	rl.last_state = rl.env.Start()
	rl.last_action = rl.agent.Start(rl.last_state)
	return array_copy(rl.last_state), array_copy(rl.last_action)
}

func (rl *RLGlue) Step() (float64, []int, []int, bool){
	this_reward, last_state, terminal := rl.env.Step(rl.last_action)
	rl.total_reward += this_reward
	rl.last_state, rl.is_terminal = array_copy(last_state), terminal

	if terminal{
		rl.num_episodes += 1
		rl.agent.End(this_reward)
		return this_reward, array_copy(rl.last_state), array_copy(rl.last_action), rl.is_terminal
	}else{
		rl.num_steps += 1
		rl.last_action = rl.agent.Step(this_reward, last_state)
		return this_reward, array_copy(rl.last_state), array_copy(rl.last_action), rl.is_terminal
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
	// is_terminal := false
	rl.Start()
	for !rl.is_terminal && (max_steps_this_episode == 0 || rl.num_steps < max_steps_this_episode){
		// _, _, _, is_terminal = rl.Step()
		rl.Step()
	}
	return rl.is_terminal
}

func (rl *RLGlue) Return() float64{
	return rl.total_reward
}

func (rl *RLGlue) Num_steps() int{
	return rl.num_steps
}

func (rl *RLGlue) Num_episodes() int{
	return rl.num_episodes
}

func NewRLGlue(agent Agent, env Environment) *RLGlue{
	return &RLGlue{agent: agent, env: env}
}