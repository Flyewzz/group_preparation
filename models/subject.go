package models

type Subject struct {
	Id           int    `json:"id"`
	UniversityId int    `json:"university_id"`
	Name         string `json:"name"`
	Semester     string `json:"semester"`
}
