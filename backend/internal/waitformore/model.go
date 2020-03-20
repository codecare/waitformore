package waitformore

import "errors"

type QueueDefinition struct {
	UserKey    string
	AdminKey   string
	QueueTitle string
	QueueState QueueState
}

type QueueState struct {
	LastPulled int
	LastCalled int
}

// this needs to be synchronised on the outside
func (state QueueState) Pull() (int, error) {
	state.LastPulled += 1
	return state.LastPulled, nil
}

// this needs to be synchronised on the outside
func (state QueueState) Call() (int, error) {
	if state.CanCall() {
		state.LastCalled += 1
		return state.LastCalled, nil
	}
	return 0, errors.New("nothing left to call")
}

func (state QueueState) CanCall() bool {
	return state.LastCalled < state.LastPulled
}

type UserView struct {
	UserKey      string
	QueueTitle   string
	QueueState   QueueState
	NumberPulled int
}

type QueueingStatus struct {
	TotalQueues int
	UpdatesLastDay int
}

