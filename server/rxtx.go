package server

import "sync"

const (
	inactive = iota
	running  = iota
	stopping = iota
	done     = iota
)

type RxTxState struct {
	state int
	lock  sync.RWMutex
}

func (s *RxTxState) Stopping() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.state == stopping
}

func (s *RxTxState) SetStop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.state = stopping
}

func (s *RxTxState) Running() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.state == running
}

func (s *RxTxState) SetRun() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.state = running
}

func (s *RxTxState) Done() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.state == done
}

func (s *RxTxState) SetDone() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.state = done
}

func (s *RxTxState) SetInactive() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.state = inactive
}

func (s *RxTxState) Inactive() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.state == inactive
}

func (s *RxTxState) Active() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.state == running || s.state == stopping
}

func NewRxTxState() *RxTxState {
	s := RxTxState{}
	s.SetInactive()
	return &s
}
