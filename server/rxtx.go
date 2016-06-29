package server

import (
	"sync"
	"time"
)

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

func (s *RxTxState) WaitDone(timeout uint32) error {
	if s.Inactive() {
		return NewError("receiver/transmitter is inactive: nothing to wait for")
	}
	start := time.Now()
	t := time.Millisecond * time.Duration(timeout)
	for time.Now().Sub(start) < t || timeout == 0 {
		if s.Done() {
			return nil
		}
		time.Sleep(time.Millisecond * 50)
	}
	return NewError("receiver/transmitter did not finish")
}
