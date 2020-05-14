package worker

import (
	"fmt"
	"sync"
)

type Worker struct {
	maxParallel int
	request     Request
	hash        Hash
	urls        []string
}

type Request interface {
	Get(url string) ([]byte, error)
}

type Hash interface {
	CreateHash(data []byte) string
}

func (w *Worker) Start(urls []string, result chan<- string, errChan chan<- error) {
	var wg sync.WaitGroup
	parallelControl := make(chan int, w.maxParallel)
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		parallelControl <- 1
		go func(i int) {
			defer wg.Done()
			respPage, err := w.request.Get(urls[i])
			if err != nil {
				errChan <- fmt.Errorf("error to get url: %s", urls[i])
			}
			respHashed := w.hash.CreateHash(respPage)
			result <- createResponse(urls[i], respHashed)
			<-parallelControl
		}(i)
	}
	wg.Wait()
}

func createResponse(address, hash string) string {
	return fmt.Sprintf("address: %s - hash: %s", address, hash)
}

func New(request Request, hash Hash, maxParallel int) *Worker {
	return &Worker{
		maxParallel: maxParallel,
		request:     request,
		hash:        hash,
	}
}
