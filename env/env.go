package env

import(
	"math/rand"
	// "testing"
	// "fmt"
)

type observation struct{
	reward 	float64
	state   []int
	is_terminal bool
}

var(
	// this_reward_observation [float64, []int, bool]
	this_reward_observation observation
	num_state = 10
)

func Init(){
	local_observation := []int{}
	this_reward_observation = observation{0, local_observation, false}
	// this_reward_observation.reward = 0
	// this_reward_observation.state = local_observation
}

func Start() []int{
	return this_reward_observation.state
}

func Step(this_action []int) (float64, []int, bool){
	the_reward := rand.NormFloat64()
	return the_reward, this_reward_observation.state, false
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