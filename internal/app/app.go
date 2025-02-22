package app

import (
	"github.com/mikolabarkouski/calculator/internal/repo"
)

type App struct {
	repo *repo.Repository
}

func NewApp(r *repo.Repository) *App {
	return &App{repo: r}
}

func (a *App) GetPackages() []int {
	packages := a.repo.GetPackages()
	return packages
}

func (a *App) AddPackage(packageSize int) {
	a.repo.AddPackage(packageSize)
}

func (a *App) DeletePackage(id string) {
	a.repo.DeletePackageById(id)
}
