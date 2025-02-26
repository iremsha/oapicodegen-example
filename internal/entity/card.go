package entity

type CardRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Cvv  int    `json:"cvv"`
}

type CardResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Cvv  int    `json:"cvv"`
}
