package models

type FunctionNode struct {
	FunctionId    int         `json:"id"`
	FunctionName  string      `json:"name"`
	FunctionPath  string      `json:"path"`
	FunctionChild interface{} `json:"children"`
}
