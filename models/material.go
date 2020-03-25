package models

type Material struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SubjectId int    `json:"subject_id"`
	TypeId    int    `json:"type_id"`
	AuthorId  int    `json:"author_id"`
	Date      string `json:"date"`
	Status    bool   `json:"-"`
}
