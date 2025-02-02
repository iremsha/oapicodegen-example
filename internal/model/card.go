package model

type Card struct {
	ID         int64  `json:"id"`
	Bank       string `json:"bank"`
	HolderName string `json:"holderName"`
}
