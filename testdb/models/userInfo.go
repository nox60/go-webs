package models

type SimpleCode struct {
	ResultCode   int    `json:"resultCode"  binding:"required"`
	Msg          string `json:"msg" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
	Name         string `json:"name" binding:"required"`
}
