package handlers

import "github.com/Flyewzz/group_preparation/interfaces"

type HandlerData struct {
	UniversityController   interfaces.UniversityController
	SubjectController      interfaces.SubjectController
	MaterialController     interfaces.MaterialController
	MaterialFileController interfaces.MaterialFileController
}

func NewHandlerData(uc interfaces.UniversityController,
	sc interfaces.SubjectController,
	mc interfaces.MaterialController, mfc interfaces.MaterialFileController) *HandlerData {
	return &HandlerData{
		UniversityController:   uc,
		SubjectController:      sc,
		MaterialController:     mc,
		MaterialFileController: mfc,
	}
}
