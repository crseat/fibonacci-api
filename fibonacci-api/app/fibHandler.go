package app

import (
	"fibonacci-api/service"
	"net/http"
	"sync"
	"time"
)

type Handler struct {
	fibService service.FibService
}

// NewSequence takes in the ResponseWriter the Request, the start time, and a pointer to the wait group. Validates the
// algorithm string and then sends it on to the fibonacci service to process.
func (ph Handler) NewSequence(w http.ResponseWriter, r *http.Request, startTime time.Time, wg *sync.WaitGroup, algo string) {
	panic("implement me")
}
