package scheduler

import "ocrawler/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	s.workerChan <- request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}


func (s *SimpleScheduler) WorkerReady(worker chan engine.Request) {
}


func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
