package worker_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/danielfmelo/myhttphash/worker"
)

type resquetMock struct {
	result []byte
	err    error
	url    string
}

func (rm *resquetMock) Get(url string) ([]byte, error) {
	rm.url = url
	return rm.result, rm.err
}

type hashMock struct {
	data   []byte
	hashed string
}

func (hm *hashMock) CreateHash(data []byte) string {
	hm.data = data
	return hm.hashed
}

type fixture struct {
	rm *resquetMock
	hm *hashMock
	w  *worker.Worker
}

func setup(maxParallel int) *fixture {
	hm := &hashMock{}
	rm := &resquetMock{}
	return &fixture{
		hm: hm,
		rm: rm,
		w: worker.New(
			rm,
			hm,
			maxParallel,
		),
	}
}

func TestStartShouldWorkAndReturnResult(t *testing.T) {
	maxParallel := 1
	f := setup(maxParallel)
	urls := []string{"test", "aaaa", "bbbb"}
	respChan := make(chan string, 10)
	errChan := make(chan error, 10)
	hashExpected := "result hashed"
	f.hm.hashed = hashExpected
	f.w.Start(urls, respChan, errChan)
	count := 0
	for {
		select {
		case resp := <-respChan:
			responseExpected := createResponse(string(urls[count]), hashExpected)
			if responseExpected != resp {
				t.Errorf("error: {%s} should be equal {%s}", responseExpected, resp)
			}
			count++

		}
		if count >= len(urls) {
			return
		}
	}

}

func TestStartShouldReturnError(t *testing.T) {
	maxParallel := 1
	f := setup(maxParallel)
	urls := []string{"test", "aaaa", "bbbb"}
	respChan := make(chan string, 10)
	errChan := make(chan error, 10)
	f.rm.err = errors.New("some error happened")
	f.w.Start(urls, respChan, errChan)
	count := 0
	for {
		select {
		case errResp := <-errChan:
			errExpected := fmt.Errorf("error to get url: %s", urls[count])
			if errResp.Error() != errExpected.Error() {
				t.Errorf("expected error: {%v} got error: {%v}", errExpected, errResp)
			}
			count++
		}
		if count >= len(urls) {
			return
		}
	}

}

func createResponse(address, hash string) string {
	return fmt.Sprintf("address: %s - hash: %s", address, hash)
}
