package models

type FetchDataRequestBody struct {
	Page  int    `json:"page"  `
	Limit int    `json:"limit" `
	Title string `json:"title"`
	Sort  string `json:"sort"`
}

type DataResponseBody struct {
	ItemId    int    `json:"itemId"  `
	ItemTilte string `json:"itemTilte"  `
	ItemPrice string `json:"itemPrice"  `
	ItemDesc  string `json:"itemDesc"  `
}
