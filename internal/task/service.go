package task

import (
	"errors"
	"time"
)

type Repository interface {
	GetAll() ([]Task, error)
	SaveAll([]Task) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddTask(title, description string, dueDate time.Time) (Task, error) {
	if title == "" {
		return Task{}, errors.New("title tidak boleh kosong!")
	}

	tasks, err := s.repo.GetAll()
	if err != nil {
		return Task{}, err
	}

	newID := s.generateID(tasks)

	task := Task{
		ID:          newID,
		Title:       title,
		Description: description,
		IsCompleted: false,
		DueDate:     dueDate,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	if err := s.repo.SaveAll(tasks); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *Service) ListTask() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) MarkDone(id int) error {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].IsCompleted = true
			return s.repo.SaveAll(tasks)
		}
	}

	return errors.New("task dengan ID tersebut tidak ditemukan!")
}

func (s *Service) generateID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func (s *Service) UpdateTask(id int, title, description string, dueDate time.Time) (Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return Task{}, err
	}

	for i, t := range tasks {
		if t.ID == id {
			if title != "" {
				tasks[i].Title = title
			}
			if description != "" {
				tasks[i].Description = description
			}
			if !dueDate.IsZero() {
				tasks[i].DueDate = dueDate
			}

			if err := s.repo.SaveAll(tasks); err != nil {
				return Task{}, err
			}
			return tasks[i], nil
		}
	}

	return Task{}, errors.New("task tidak ditemukan")
}

func (s *Service) DeleteTask(id int) error {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return s.repo.SaveAll(tasks)
		}
	}

	return errors.New("task tidak ditemukan")
}
