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

type Environment interface {
	Init()
	Start() []int
	Step([]int) (float64, []int, bool)
	Cleanup()
	Message(string) interface{}
}

type RLGlue struct {
	env         Environment
	agent       Agent
	totalReward float64
	numSteps    int
	numEpisodes int
	lastState   []int
	lastAction  []int
	isTerminal  bool
}

func array_copy(arr []int) []int {
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	return cpy
}

// func array_copy(arr []reflect.Type){
// 	cpy := reflect.ArrayOf(len(arr), reflect.TypeOf(arr))
// 	reflect.Copy(cpy, arr)
// 	return cpy
// }

func (rl *RLGlue) Init() {
	rl.env.Init()
	rl.agent.Init()
	rl.totalReward = 0.0
	rl.numSteps = 0
	rl.numEpisodes = 0
	rl.isTerminal = false
}

func (rl *RLGlue) Start() ([]int, []int) {
	rl.totalReward = 0.0
	rl.isTerminal = false
	rl.numSteps = 1
	rl.lastState = rl.env.Start()
	rl.lastAction = rl.agent.Start(rl.lastState)
	return array_copy(rl.lastState), array_copy(rl.lastAction)
}

func (rl *RLGlue) Step() (float64, []int, []int, bool) {
	thisReward, lastState, terminal := rl.env.Step(rl.lastAction)
	rl.totalReward += thisReward
	rl.lastState, rl.isTerminal = array_copy(lastState), terminal

	if terminal {
		rl.numEpisodes++
		rl.agent.End(thisReward)

	} else {
		rl.numSteps++
		rl.lastAction = rl.agent.Step(thisReward, lastState)
	}
	return thisReward, array_copy(rl.lastState), array_copy(rl.lastAction), rl.isTerminal
}

func (rl *RLGlue) Cleanup() {
	rl.env.Cleanup()
	rl.agent.Cleanup()
}

// TODO
func (rl *RLGlue) Message(inMessage string) interface{} {
	return nil
}

func (rl *RLGlue) Episode(maxStepsThisEpisode int) bool {
	// is_terminal := false
	rl.Start()
	for !rl.isTerminal && (maxStepsThisEpisode == 0 || rl.numSteps < maxStepsThisEpisode) {
		// _, _, _, is_terminal = rl.Step()
		rl.Step()
	}
	return rl.isTerminal
}

func (rl *RLGlue) Return() float64 {
	return rl.totalReward
}

func (rl *RLGlue) Num_steps() int {
	return rl.numSteps
}

func (rl *RLGlue) Num_episodes() int {
	return rl.numEpisodes
}

func NewRLGlue(agent Agent, env Environment) *RLGlue {
	return &RLGlue{agent: agent, env: env}
}
