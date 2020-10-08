package models

type UserInfo struct {
	Code         int    `json:"code" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Roles        string `json:"roles" binding:"required"`
	Page         int    `json:"page"  `
	Limit        int    `json:"limit" `
	ForCount     bool
}

type User struct {
	AccountId int    `json:"accountId"`
	UserName  string `json:"userName"`
	RealName  string `json:"realName"`
}

func (reqBody *UserInfo) GetStartByPageAndLimit() int {
	result := (reqBody.Page - 1) * reqBody.Limit
	return result
}
