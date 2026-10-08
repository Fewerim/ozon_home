package filerepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	"os"
	"path/filepath"
	"strings"
)

// Для упрощения специально не использовал конфиг, по хорошему бы конечно, вынести все константы туда
const (
	defaultStoragePath     = "./data"
	defaultNameStorageFile = "data.json"
	defaultNextID          = 1
)

type FileRepository struct {
	path   string
	nextID domain.TaskID
	tasks  map[domain.TaskID]*domain.Task
}

// NewFileRepository - создает репозиторий для хранения задач в файле
func NewFileRepository(storagePath string, nameFile string) (*FileRepository, error) {
	var path string
	const ext = ".json"

	if strings.TrimSpace(storagePath) == "" || strings.TrimSpace(nameFile) == "" {
		path = filepath.Join(
			defaultStoragePath,
			defaultNameStorageFile,
		)
	} else {
		path = filepath.Join(storagePath, nameFile+ext)
	}

	repo := &FileRepository{
		path:   path,
		nextID: defaultNextID,
		tasks:  make(map[domain.TaskID]*domain.Task),
	}

	if err := repo.load(); err != nil {
		return nil, fmt.Errorf("failed to load file for storage: %w", err)
	}

	return repo, nil
}

func (r *FileRepository) updateNextID() { r.nextID++ }

// load - загружает задачи в оперативную память из файла
func (r *FileRepository) load() error {
	data, err := os.ReadFile(r.path)
	if errors.Is(err, os.ErrNotExist) {
		// файл будет создаваться при создании первой задачи или при импорте файла
		return nil
	}
	if err != nil {
		return fmt.Errorf("read storage file: %w", err)
	}

	// десериализуем данные из файла в модель
	var model TasksModel
	if err := json.Unmarshal(data, &model); err != nil {
		return fmt.Errorf("decode storage file: %w", err)
	}

	// преобразуем модель в бизнес сущность и получаем следующий айди для задачи
	tasks, nextID := model.ToDomain()
	if nextID < defaultNextID {
		return fmt.Errorf("invalid next_id %d in storage file", nextID)
	}

	loaded := make(map[domain.TaskID]*domain.Task, len(tasks))
	for i := range tasks {
		curTask := &tasks[i]
		if curTask.ID <= 0 || curTask.ID >= nextID {
			return fmt.Errorf("invalid task ID %d in storage file", curTask.ID)
		}
		if _, exists := loaded[curTask.ID]; exists {
			return fmt.Errorf("duplicate task ID %d in storage file", curTask.ID)
		}

		// проверяем, что задачи, лежащие в файле валидны, и соответсвуют БП
		if err := curTask.Validate(); err != nil {
			return fmt.Errorf("invalid stored task %d: %w", curTask.ID, err)
		}
		loaded[curTask.ID] = curTask
	}

	r.tasks = loaded
	r.nextID = nextID

	return nil
}

// save - сохранить задачи в файл
func (r *FileRepository) save() error {
	// получаем все задачи из оперативной памяти и преобразуем их в модель
	model := NewTasksModel(r.getList(), r.nextID)

	// сериализуем данные в json
	data, err := json.MarshalIndent(model, "", "	")
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}

	// если нет директории - создаем, если есть, то идем дальше
	if err := os.MkdirAll(filepath.Dir(r.path), 0755); err != nil {
		return fmt.Errorf("create storage directory: %w", err)
	}

	// записываем в файл
	if err := os.WriteFile(r.path, data, 0644); err != nil {
		return fmt.Errorf("write storage file: %w", err)
	}

	return nil
}
