package models

type TokenPayload struct {
	AccountId int    `json:"accountId"`
	MenuItems string `json: "menuItems"`
	Anothers  string `json:"anoItems"`
}
