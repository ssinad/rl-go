package main

import (
	"fmt"
)

var (
	numEpisodes = 8000
	maxSteps    = 10000
	done        = make(chan bool)
)

// Maybe this should be moved to inside RLGlue itself
func singleRun(run int) {
	fmt.Printf("run number %d\n\n", (run + 1))
	numStates := 100
	agent := NewMCAgent(numStates, 0.9, 0.5)
	rl := NewRLGlue(agent, NewToyEnvironment(numStates))
	rl.Init()

	for episode := 0; episode < numEpisodes; episode++ {
		rl.Episode(maxSteps)
	}
	rl.Cleanup()
	fmt.Print("[ ")
	for i := 0; i < numStates; i++ {
		fmt.Println(agent.V.At(i, 0), ", ")
	}
	fmt.Println(" ]")

	done <- true
}

func main() {
	numRuns := 1

	for run := 0; run < numRuns; run++ {
		go singleRun(run)
	}
	for run := 0; run < numRuns; run++ {
		<-done
	}
}
