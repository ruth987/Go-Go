package data

import (
	"errors"
	"project_5/models"
	"time"
)

type TaskService struct {
	tasks []models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: []models.Task{
			{
				ID:          "1",
				Title:       "Task Manager Project",
				Description: "Add/View/Delete Tasks",
				DueDate:     time.Now(),
				Status:      "In Progress",
			},
			{
				ID:          "2",
				Title:       "Books Management Project",
				Description: "Add/View/Delete Books",
				DueDate:     time.Now().AddDate(0, 0, -1),
				Status:      "Completed",
			},
		},
	}
}

func (s *TaskService) GetAllTasks() []models.Task {
	return s.tasks
}

func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) CreateTask(task models.Task) error {
	s.tasks = append(s.tasks, task)
	return nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range s.tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				s.tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				s.tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				s.tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				s.tasks[i].Status = updatedTask.Status
			}
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) DeleteTask(id string) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
