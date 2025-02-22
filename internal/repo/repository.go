package repo

import "log"

type Repository struct{}

func NewRepository() *Repository {
	log.Println("Initializing repository")
	return &Repository{}
}
