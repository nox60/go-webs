package models

type FunctionNode struct {
	FunctionId       int         `json:"id"`
	Number           int         `json:"number"`
	Order            int         `json:"order"`
	Name             string      `json:"name"`
	Path             string      `json:"path"`
	ParentFunctionId int         `json:"parentId"`
	HasChildren      bool        `json:"hasChildren"`
	Leaf             bool        `json:"leaf"`
	Parents          interface{} `json:"parents"`
}
