package entity

type BankRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BankResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
