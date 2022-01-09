package scheduler

import (
	"crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	// 每个worker都有自己的chan
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) Run() {

	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			// 这里不直接把request给到wrokerChan，是想所有chan操作都放在select中
			var activeRequest engine.Request
			// chan 是 nil 是不会被select到的
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerChain() chan engine.Request {
	return make(chan engine.Request)
}
