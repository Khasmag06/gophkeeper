package models

import "time"

type Record struct {
	EncryptedData string `json:"encrypted_data"`
}

type LoginCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TextData struct {
	Data string `json:"data"`
}

type BinaryData struct {
	Data []byte `json:"data"`
}

type BankCard struct {
	CardNumber     string    `json:"card_number"`
	ExpirationDate time.Time `json:"expiration_date"`
	CVV            string    `json:"cvv"`
}
