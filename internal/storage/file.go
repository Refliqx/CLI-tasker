package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Refliqx/tasker-project/internal/task"
)

type FileRepository struct {
	path string
}

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{
		path: path,
	}
}

func (r *FileRepository) GetAll() ([]task.Task, error) {
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return []task.Task{}, nil
	}

	data, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *FileRepository) SaveAll(tasks []task.Task) error {
	if err := os.MkdirAll(filepath.Dir(r.path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.path, data, 0644)
}
