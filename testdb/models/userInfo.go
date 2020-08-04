package models

type UserInfo struct {
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
	Name         string `json:"name" binding:"required"`
}
