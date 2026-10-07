package filter

import (
	"cmp"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	"slices"
	"time"
)

// FilterAndSort - фильтр для получения списка задач и сортировка их по дедлайну/айди
// Следующий фильтр: статус, только просроченные
func FilterAndSort(tasks []domain.Task, targetStatus domain.TaskStatus, exited bool) []domain.Task {
	result := make([]domain.Task, 0, len(tasks))
	now := time.Now()

	for _, task := range tasks {
		if targetStatus != "" && task.Status != targetStatus {
			continue
		}

		if exited && (!task.Deadline.Before(now) ||
			task.Status == domain.StatusDone || task.Status == domain.StatusCanceled) {
			continue
		}

		result = append(result, task)
	}

	sort(result)

	return result
}

func sort(tasks []domain.Task) {
	slices.SortFunc(tasks, func(t1, t2 domain.Task) int {
		if n := t1.Deadline.Compare(t2.Deadline); n != 0 {
			return n
		}

		return cmp.Compare(t1.ID, t2.ID)
	})
}
