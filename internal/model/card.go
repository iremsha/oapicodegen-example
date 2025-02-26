package model

type Card struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Cvv  int    `json:"cvv"`
}
