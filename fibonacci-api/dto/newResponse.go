package dto

import "time"

type NewResponse struct {
	Input    int
	Fib      int64
	Duration time.Duration
	Algo     string
	Status   string
}
