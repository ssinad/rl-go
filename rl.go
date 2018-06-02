package main

// import(
// 	"reflect"
// )

//Agent Any agent should comply to this interface
type Agent interface {
	Init()
	Start([]int) []int
	Step(float64, []int) []int
	End(float64)
	Cleanup()
	Message(string) interface{}
}

//Environment Any environment should comply to this interface
type Environment interface {
	Init()
	Start() []int
	Step([]int) (float64, []int, bool)
	Cleanup()
	Message(string) interface{}
}

//RLGlue Mediator between the agent and the environment
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

func arrayCopy(arr []int) []int {
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
	return arrayCopy(rl.lastState), arrayCopy(rl.lastAction)
}

func (rl *RLGlue) Step() (float64, []int, []int, bool) {
	thisReward, lastState, terminal := rl.env.Step(rl.lastAction)
	rl.totalReward += thisReward
	rl.lastState, rl.isTerminal = arrayCopy(lastState), terminal

	if terminal {
		rl.numEpisodes++
		rl.agent.End(thisReward)

	} else {
		rl.numSteps++
		rl.lastAction = rl.agent.Step(thisReward, lastState)
	}
	return thisReward, arrayCopy(rl.lastState), arrayCopy(rl.lastAction), rl.isTerminal
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

func (rl *RLGlue) NumSteps() int {
	return rl.numSteps
}

func (rl *RLGlue) NumEpisodes() int {
	return rl.numEpisodes
}

func NewRLGlue(agent Agent, env Environment) *RLGlue {
	return &RLGlue{agent: agent, env: env}
}
