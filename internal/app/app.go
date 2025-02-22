package app

import "github.com/mikolabarkouski/calculator/internal/repo"

type App struct {
	repo *repo.Repository
}

func NewApp(r *repo.Repository) *App {
	return &App{repo: r}
}
