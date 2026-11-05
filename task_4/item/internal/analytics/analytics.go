package analytics

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Record struct {
	Operation string
	Duration  time.Duration
}

const (
	batchSize    = 5
	flushTimeout = 5 * time.Second
)

type Analytics struct {
	records chan Record
	total   atomic.Int64

	mu    sync.Mutex
	saved []Record
}

func New() *Analytics {
	return &Analytics{
		records: make(chan Record, batchSize*4),
	}
}

func (a *Analytics) Log(operation string, d time.Duration) {
	if a == nil {
		return
	}
	a.records <- Record{Operation: operation, Duration: d}
}

func (a *Analytics) Run(ctx context.Context) {
	ticker := time.NewTicker(flushTimeout)
	defer ticker.Stop()

	batch := make([]Record, 0, batchSize)
	flush := func() {
		if len(batch) == 0 {
			return
		}
		a.store(batch)
		batch = make([]Record, 0, batchSize)
	}

	for {
		select {
		case <-ctx.Done():
			flush()
			return
		case r := <-a.records:
			a.total.Add(1)
			batch = append(batch, r)
			if len(batch) == batchSize {
				flush()
			}
		case <-ticker.C:
			flush()
		}
	}
}

func (a *Analytics) store(batch []Record) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.saved = append(a.saved, batch...)
	log.Printf("analytics: stored batch of %d records (total %d)", len(batch), a.total.Load())
}
