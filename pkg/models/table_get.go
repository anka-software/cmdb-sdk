package models

type GetTableItem struct {
	Result []GetTable
}

type GetTable struct {
	SysClassName string `json:"sys_class_name"`
	// Probably Unique Identifier.
	SysId string `json:"sys_id"`
	Name  string `json:"name,omitempty"`
}
