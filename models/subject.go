package models

type Subject struct {
	Id           int    `json:"id"`
	UniversityId int    `json:"university_id,omitempty"`
	Name         string `json:"name"`
	Semester     string `json:"semester"`
}
