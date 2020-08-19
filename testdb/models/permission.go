package models

type FunctionNode struct {
	FunctionId   int    `json:"id"`
	FunctionName string `json:"name"`
	FunctionPath string `json:"path"`
	ParentId     int    `json:"parentId"`
	HasChildren  bool   `json:"hasChildren"`
}
