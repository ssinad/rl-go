package rl

import(
	"rl/env"
	"rl/agent"
	// "testing"
	// "fmt"
)

var(
	// this_reward_observation [float64, []int, bool]
	// this_reward_observation env.Observation
	last_action []int
	total_reward float64
	num_steps int
	num_episodes int
	// num_state = 10
)

func RLGlue(){
	// local_observation := []int{}
	// this_reward_observation = Observation{0, local_observation, false}
	// this_reward_observation.reward = 0
	// this_reward_observation.state = local_observation
}

func Init(){
	env.Init()
	agent.Init()
	total_reward = 0.0
	num_steps = 0
	num_episodes = 0
}

func Start() ([]int, []int){
	total_reward = 0.0
	num_steps = 1
	last_state := env.Start()
	last_action := agent.Start(last_state)
	return last_state, last_action
}

func Step() (float64, []int, []int, bool){
	this_reward, last_state, terminal := env.Step(last_action)
	// this_reward := observation.Reward
	// last_state := observation.State
	// terminal := observation.Is_terminal 
	total_reward += this_reward

	if terminal{
		num_episodes += 1
		agent.End(this_reward)
		return this_reward, last_state, last_action, terminal
	}else{
		num_steps += 1
		last_action = agent.Step(this_reward, last_state)
		return this_reward, last_state, last_action, terminal
	}
}

func Cleanup(){
	env.Cleanup()
	agent.Cleanup()
}

// TODO
func Message(inMessage string){

}

func Episode(max_steps_this_episode int) bool{
	is_terminal := false
	Start()
	for !is_terminal && (max_steps_this_episode == 0 || num_steps < max_steps_this_episode){
		_, _, _, is_terminal = Step()
	}
	return is_terminal
}
// func main(){
// 	Init()
// 	fmt.Println(last_action)

// 	tmp := Start([]int{1})
// 	fmt.Println(last_action)
// 	fmt.Println(tmp)

// 	for i := 0; i < 10; i++ {
// 		tmp := Step(2, []int{1})
// 		fmt.Println(last_action)
// 		fmt.Println(tmp)
// 	}

// 	End(3)
// 	fmt.Println(last_action)
// }