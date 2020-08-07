package models

type FetchDataRequestBody struct {
	Page     int    `json:"page"  `
	Limit    int    `json:"limit" `
	Title    string `json:"title"`
	Sort     string `json:"sort"`
	ForCount bool
}

func (reqBody *FetchDataRequestBody) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}

type DataResponseBody struct {
	ItemId    int    `json:"itemId"  `
	ItemTitle string `json:"itemTitle"  `
	ItemPrice string `json:"itemPrice"  `
	ItemDesc  string `json:"itemDesc"  `
}
