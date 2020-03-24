package handlers

import "github.com/Flyewzz/group_preparation/interfaces"

type HandlerData struct {
	UniversityController interfaces.UniversityController
	SubjectController    interfaces.SubjectController
	MaterialController   interfaces.MaterialController
}

func NewHandlerData(uc interfaces.UniversityController,
	sc interfaces.SubjectController, mc interfaces.MaterialController) *HandlerData {
	return &HandlerData{
		UniversityController: uc,
		SubjectController:    sc,
		MaterialController:   mc,
	}
}
