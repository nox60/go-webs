package models

type HttpResult struct {
	Code  int         `json:"code"  binding:"required"`
	Msg   string      `json:"msg" binding:"required"`
	Token string      `json:"token" binding:"required"`
	Data  interface{} `json:"data" binding:"required"`
}
