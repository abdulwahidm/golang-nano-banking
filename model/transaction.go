package model

type Transaction struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}
