package app

import (
	"fibonacci-api/dto"
	"fibonacci-api/service"
	"net/http"
	"sync"
)

type fibHandler struct {
	fibService service.FibService
}

// NewSequence takes in the ResponseWriter the Request, the start time, and a pointer to the wait group. Validates the
// algorithm string and then sends it on to the fibonacci service to process.
func (fh fibHandler) NewSequence(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup, algo string) {
	var request = dto.NewRequest{}

	//Build the request object
	err := r.ParseForm()
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	}
	//Get a validate algo
	request.Algorithm = algo
	appError := request.ValidateAlgo()
	if appError != nil {
		writeResponse(w, http.StatusBadRequest, appError.AsMessage())
		return
	}
	//Get a validate inputNum
	inputNum := r.Form.Get("input")
	num, appError := request.ValidateInputNum(inputNum)
	if appError != nil {
		writeResponse(w, http.StatusBadRequest, appError.AsMessage())
		return
	}
	request.Input = num
	//Process Input
	response, appError := fh.fibService.NewSequence(request, wg)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, response.Id)
}

// FindBy takes in the ResponseWriter and the fib identifier. Validates the identifier then passes it on to the
// fibService to process
func (fh fibHandler) FindBy(w http.ResponseWriter, id string) {
	var request = dto.NewRequest{}
	fibId, appError := request.ValidateId(id)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	request.Id = fibId
	sequence, appError := fh.fibService.FindById(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, sequence)
	}
}
