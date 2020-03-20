package waitformore

import (
	"fmt"
	"testing"
)

func Test_iterateQueues(t *testing.T) {

	queueStates := []*QueueState{{
		LastPulled: 1,
		LastCalled: 3,
	}, {
		LastPulled: 4,
		LastCalled: 6,
	}}

	fmt.Printf("queue states: %v\n", queueStates)

	for i, state := range queueStates {
		queueStates[0].LastPulled = 1111111111
		state.LastCalled = 33333
		fmt.Printf("queue state: %d %v\n", i, state)
	}

	fmt.Printf("queue states: %v\n", queueStates)
}
