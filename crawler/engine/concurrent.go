package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}


func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i:=0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicateUrl(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <- out
		for _, item := range result.Items {
			go func(item Item) {
				e.ItemChan <- item
			}(item)
		}

		// URL dedup
		for _, request := range result.Requests {
			if isDuplicateUrl(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			go func() {
				out <- result
			}()
		}
	}()
}

var visitedUrls  = make(map[string]bool)
func isDuplicateUrl(url string) bool {
	if _, ok := visitedUrls[url]; ok {
		return true
	}else{
		visitedUrls[url] = true
		return false
	}
}
