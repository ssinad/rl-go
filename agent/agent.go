package agent

import(
	"math/rand"
	// "testing"
	// "fmt"
)

// type Agent struct{
// 	value_function []float64
// }

var(
	last_action []int
	num_actions = 10
	Value_function = []float64{0.1, 0.2, 0.3}
)

func Init(){
	last_action = []int{0}
}

func Start(this_observation []int) []int{
	last_action[0] = rand.Intn(num_actions)
	local_action := []int{0}
	local_action[0] = rand.Intn(num_actions)
	return local_action
}

func Step(reward float64, this_observation []int) []int{
	local_action := []int{0}
	local_action[0] = rand.Intn(num_actions)
	last_action = local_action
	return last_action
}

func End(reward float64){
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