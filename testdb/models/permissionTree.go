package models

type PermissionNode struct {
	PermissionId    int         `json:"id"`
	PermissionName  string      `json:"name"`
	PermissionPath  string      `json:"path"`
	PermissionChild interface{} `json:"child"`
}
