package usecases

import "github.com/jonsampson/bubbletea-playground/internal/domain"

type CreateProjectUseCase struct {
}

func NewCreateProjectUseCase() *CreateProjectUseCase {
	return &CreateProjectUseCase{}
}

func (u CreateProjectUseCase) CreateProject(validBubbleTeaPlayground domain.ValidBubbleTeaPlayground) error {
	return nil
}
