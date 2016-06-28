package server

type IOState struct {
	Started  bool
	Stopping bool
	Done     bool
}

var NotStarted = IOState{}
var Started = IOState{Started: true}
var Stopping = IOState{Stopping: true}
var Done = IOState{Done: true}

type NetIO interface {
	Stop()
	Start()
	Join(timeout uint32) error
	State() IOState
	SetState(newState IOState)
	Stats() (Stats, error)
}

type Stats struct {
	Received uint32
	KDropped uint32
	IDropped uint32
}
