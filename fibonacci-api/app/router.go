package app

import (
	"encoding/json"
	"fibonacci-api/errs"
	"fibonacci-api/logger"
	"net/http"
	"path"
	"strings"
	"sync"
)

type Router struct {
	Handler   *fibHandler
	WaitGroup *sync.WaitGroup
	quitChan  *chan bool
}

//define routes
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string

	//Check for invalid method.
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		invalidMethodError(w)
	}

	//get endpoint
	head, r.URL.Path = shiftPath(r.URL.Path)

	//Define routes.
	switch head {
	case "fib":
		//check for algo
		var algorithm string
		algorithm, r.URL.Path = shiftPath(r.URL.Path)
		router.Handler.NewSequence(w, r, router.WaitGroup, algorithm)
	case "find":
		//check for id
		var id string
		id, r.URL.Path = shiftPath(r.URL.Path)
		router.Handler.FindBy(w, id)
	case "shutdown":
		//Shutdown gracefully
		shutdown(*router.quitChan)
	default:
		logger.DebugLogger.Println("Attempted invalid endpoint = ", head)
		invalidEndpointError(w)
	}
}

// shiftPath splits the given path into the first segment (head) and  the rest (tail).
// For example, "/foo/bar/baz" gives "foo", "/bar/baz".
//source: https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func shutdown(quit chan bool) {
	quit <- true
}

func invalidMethodError(w http.ResponseWriter) {
	appError := errs.NewValidationError("Method is not supported")
	writeResponse(w, http.StatusNotFound, appError.AsMessage())
}

func invalidEndpointError(w http.ResponseWriter) {
	appError := errs.NewValidationError("Please provide a valid endpoint")
	writeResponse(w, http.StatusNotFound, appError.AsMessage())
}

// writeResponse formats all http responses to client into json
func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	// We need to define the header here or the json/xml response will come across as plain text
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
