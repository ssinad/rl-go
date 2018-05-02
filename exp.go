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
	// rl := RLGlue{agent: &SimpleAgent{}, env: &SimpleEnvironment{}}
	rl := NewRLGlue(NewSimpleAgent(), NewSimpleEnvironment())
	rl.Init()

	for episode := 0; episode < numEpisodes; episode++ {
		rl.Episode(maxSteps)
	}
	rl.Cleanup()
	done <- true
}

func main() {

	numRuns := 10
	runs := make([]float64, numRuns)
	runs[0] = 0.1
	runs[1] = 0.2
	runs[2] = 0.3
	for _, value := range runs {
		fmt.Println(value)
	}

	for run := 0; run < numRuns; run++ {
		go singleRun(run)
	}
	for run := 0; run < numRuns; run++ {
		<-done
	}
}
