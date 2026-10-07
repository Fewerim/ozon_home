package filerepository

import (
	"github.com/Fewerim/ozon_home/task_1/internal/app/repository/filter"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
)

// GetListTasks - возвращает список задач из хранилища, который фильтруется: по статусу и просроченным заданиям. Результирующий список сортируется по дедлайну задач, если дедлайн совпадает, то по айди
func (r *FileRepository) GetListTasks(targetStatus domain.TaskStatus, exited bool) ([]domain.Task, error) {
	tasks := r.getList()

	// применение фильтра и сортировки
	return filter.FilterAndSort(tasks, targetStatus, exited), nil
}

// getList - получить список существующих задач
func (r *FileRepository) getList() []domain.Task {
	// инициализируем слайс, чтобы в него передать все значения хранящиеся в репо (защищаем от внешнего изменения задач)
	result := make([]domain.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		result = append(result, *task)
	}

	return result
}
