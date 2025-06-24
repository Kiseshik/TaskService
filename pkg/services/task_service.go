package services

import (
	"context"
	"time"

	"github.com/Kiseshik/TaskService.git/pkg/models"
	"github.com/Kiseshik/TaskService.git/pkg/storage"
	"github.com/google/uuid"
)

type TaskService struct {
	store storage.Storage
}

func NewTaskService(store storage.Storage) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) CreateTask(ctx context.Context) *models.Task {
	task := &models.Task{
		ID:        generateID(),
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	}

	s.store.Create(task)

	go s.executeTask(task)

	return task
}

func (s *TaskService) executeTask(task *models.Task) {
	s.updateTaskStatus(task, models.StatusRunning, nil)

	time.Sleep(3*time.Minute + time.Duration(task.ID[len(task.ID)-1]%2)*time.Minute)

	s.updateTaskStatus(task, models.StatusCompleted, map[string]interface{}{
		"data": "Task completed successfully",
	})
}

func (s *TaskService) updateTaskStatus(task *models.Task, status models.TaskStatus, result interface{}) {
	storedTask, exists := s.store.Get(task.ID)
	if !exists {
		return
	}

	if status == models.StatusRunning && storedTask.StartedAt == nil {
		now := time.Now()
		storedTask.StartedAt = &now
	} else if (status == models.StatusCompleted || status == models.StatusFailed) && storedTask.CompletedAt == nil {
		now := time.Now()
		storedTask.CompletedAt = &now
		if storedTask.StartedAt != nil {
			storedTask.Duration = now.Sub(*storedTask.StartedAt).Seconds()
		}
	}

	storedTask.Status = status
	storedTask.Result = result

	s.store.Update(storedTask)
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*models.Task, bool) {
	return s.store.Get(id)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) bool {
	return s.store.Delete(id)
}

func generateID() string {
	return uuid.New().String()
}
