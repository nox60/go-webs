package models

type SimpleCode struct {
	ResultCode int    `json:"resultCode"  binding:"required"`
	Msg        string `json:"msg" binding:"required"`
	Token      string `json:"token" binding:"required"`
}
