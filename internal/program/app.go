package program

import (
	"github.com/jonsampson/bubbletea-playground/internal/handlers/bubbletea"
	"github.com/jonsampson/bubbletea-playground/internal/usecases"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a App) Run() {
	createProjectUseCase := usecases.NewCreateProjectUseCase()
	bubbletea.Run(createProjectUseCase)
}
