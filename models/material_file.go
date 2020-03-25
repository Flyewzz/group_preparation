package models

type MaterialFile struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Path       string `json:"-"`
	MaterialId int    `json:"material_id,omitempty"`
}
