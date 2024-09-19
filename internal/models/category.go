package models

type Category struct {
	ID       string `json:"category_id"`
	Name     string `json:"category_name"`
	ParentId int    `json:"parent_id"`
}
