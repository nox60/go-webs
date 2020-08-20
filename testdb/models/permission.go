package models

type FunctionNode struct {
	FunctionId       int    `json:"id"`
	Number           int    `json:"number"`
	Order            int    `json:"order"`
	Name             string `json:"name"`
	Path             string `json:"path"`
	ParentFunctionId int    `json:"parentId"`
	HasChildren      bool   `json:"hasChildren"`
}

type TreeNode struct {
	Name  string `json:"name"`
	Other string `json:"testColumn"`
	Leaf  bool   `json:"leaf"`
}
