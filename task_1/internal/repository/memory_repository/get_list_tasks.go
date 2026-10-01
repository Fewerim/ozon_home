package memoryrepository

import (
	"cmp"
	"slices"
	"time"
	"todoList/internal/core/domain"
)

// GetListTasks - возвращает список задач, который фильтруется: по статусу и просроченным заданиям. Результирующий список сортируется по дедлайну, если дедлайн совпадает, то по айди
func (r *MemoryRepository) GetListTasks(targetStatus domain.TaskStatus, exited bool) ([]domain.Task, error) {
	// применение фильтров
	tasks := r.filtered(targetStatus, exited)

	// инициализируем слайс, чтобы в него передать все значения хранящиеся в репо (защищаем от внешнего изменения задач)
	result := make([]domain.Task, 0, len(tasks))

	for _, task := range tasks {
		result = append(result, *task)
	}

	// сортировка
	sorted(result)

	return result, nil
}

func (r *MemoryRepository) filtered(targetStatus domain.TaskStatus, exited bool) map[int64]*domain.Task {
	result := make(map[int64]*domain.Task, len(r.tasks))
	now := time.Now()

	for id, task := range r.tasks {
		if task.Status != targetStatus {
			continue
		}

		if exited && (!task.Deadline.Before(now) ||
			task.Status == domain.StatusDone || task.Status == domain.StatusCanceled) {
			continue
		}

		result[id] = task
	}

	return result
}

func sorted(tasks []domain.Task) {
	slices.SortFunc(tasks, func(t1, t2 domain.Task) int {
		if n := t1.Deadline.Compare(t2.Deadline); n != 0 {
			return n
		}

		return cmp.Compare(t1.ID, t2.ID)
	})
}
