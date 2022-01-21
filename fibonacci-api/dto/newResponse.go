package dto

import "math/big"

type NewResponse struct {
	Input    int     `json:"input"`
	Fib      big.Int `json:"fib"`
	Duration int64   `json:"duration"`
	Algo     string  `json:"algo"`
	Status   string  `json:"status"`
	Id       int64   `json:"id"`
}
