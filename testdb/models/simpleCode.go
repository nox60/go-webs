package models

type SimpleCode struct {
	ResultCode int    `json:"resultCode"  binding:"required"`
	Msg        string `json:"msg" binding:"required"`
}
