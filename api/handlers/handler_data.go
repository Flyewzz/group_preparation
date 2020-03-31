package handlers

import "github.com/Flyewzz/group_preparation/interfaces"

type HandlerData struct {
	UniversityController interfaces.UniversityController
	SubjectController    interfaces.SubjectController
	MaterialController   interfaces.MaterialController
	AuthController       interfaces.AuthController
}

func NewHandlerData(
	uc interfaces.UniversityController,
	sc interfaces.SubjectController,
	mc interfaces.MaterialController,
	ac interfaces.AuthController,
) *HandlerData {
	return &HandlerData{
		UniversityController: uc,
		SubjectController:    sc,
		MaterialController:   mc,
		AuthController:       ac,
	}
}
