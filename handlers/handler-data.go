package handlers

import "github.com/Flyewzz/group_preparation/interfaces"

type HandlerData struct {
	UniversityController interfaces.UniversityController
	// SubjectController    interfaces.SubjectController
}

func NewHandlerData(uc interfaces.UniversityController) *HandlerData {
	return &HandlerData{
		UniversityController: uc,
		// SubjectController:    sc,
	}
}
