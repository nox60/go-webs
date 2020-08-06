package models

type FetchDataRequestBody struct {
	page  int    `json:"page"  `
	limit int    `json:"limit" `
	title string `json:"title"`
	sort  string `json:"sort"`
}

type DataResponseBody struct {
	Id         int    `json:"id"  `
	Timestamp  string `json:"timestamp"  `
	TypeValue  string `json:"type"  `
	Author     string `json:"author"  `
	Reviewer   string `json:"reviewer"  `
	Importance string `json:"importance"  `
	Pageviews  string `json:"pageviews"  `
	Status     string `json:"status"  `
}
