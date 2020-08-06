package models

type UserInfo struct {
	Code         int    `json:"code" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Roles        string `json:"roles" binding:"required"`
}
