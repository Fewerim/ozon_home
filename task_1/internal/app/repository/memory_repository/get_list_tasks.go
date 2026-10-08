package memoryrepository

import (
	tasklist "github.com/Fewerim/ozon_home/task_1/internal/app/repository/task_list"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
)

// GetListTasks - возвращает список задач из хранилища, который фильтруется: по статусу и просроченным заданиям. Результирующий список сортируется по дедлайну задач, если дедлайн совпадает, то по айди
func (r *MemoryRepository) GetListTasks(targetStatus domain.TaskStatus, exited bool) ([]domain.Task, error) {
	// инициализируем слайс, чтобы в него передать все значения хранящиеся в репо (защищаем от внешнего изменения задач)
	result := make([]domain.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		result = append(result, *task)
	}

	// применение фильтра и сортировки
	return tasklist.FilterAndSort(result, targetStatus, exited), nil
}
