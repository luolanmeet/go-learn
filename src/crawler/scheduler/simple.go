package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChain(workerChan chan engine.Request) {
	s.workerChan = workerChan
}
