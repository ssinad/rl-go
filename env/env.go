package env

import(
	"math/rand"
	// "testing"
	// "fmt"
)

type Observation struct{
	Reward 	float64
	State   []int
	Is_terminal bool
}

var(
	// this_reward_observation [float64, []int, bool]
	this_reward_observation Observation
	num_state = 10
)

func Init(){
	local_observation := []int{}
	this_reward_observation = Observation{0, local_observation, false}
	// this_reward_observation.reward = 0
	// this_reward_observation.state = local_observation
}

func Start() []int{
	return this_reward_observation.state
}

func Step(this_action []int) Observation{
	the_reward := rand.NormFloat64()
	this_reward_observation = Observation{the_reward, this_reward_observation.state, false}
	return this_reward_observation
}

func Cleanup(){
}

// TODO
func Message(inMessage string){

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