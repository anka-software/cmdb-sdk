package models

type GetTableItem struct {
	Result []GetTable
}

type GetTable struct {
	SysClassName string `json:"sys_class_name,omitempty"`
	Name         string `json:"name,omitempty"`
}
