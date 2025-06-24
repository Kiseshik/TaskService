package storage

import "github.com/Kiseshik/TaskService.git/pkg/models"

type Storage interface {
	Create(task *models.Task)
	Get(id string) (*models.Task, bool)
	Update(task *models.Task)
	Delete(id string) bool
}
