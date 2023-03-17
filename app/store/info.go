package store

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Info struct {
	UUID           uuid.UUID
	GenerationTime time.Time
	Timeout        time.Duration
	mux            sync.RWMutex
}

func NewInfo(ctx context.Context, t time.Duration) *Info {
	info := &Info{
		Timeout: t,
	}

	info.GenerateUUID()
	info.regenerate(ctx)

	return info
}

func (i *Info) regenerate(ctx context.Context) {
	go func() {
		for {
			select {
			case <-time.After(i.Timeout):
				i.GenerateUUID()
			case <-ctx.Done():
				log.Println("Finish regenerate")
				return
			}
		}
	}()
}

func (i *Info) GenerateUUID() {
	i.mux.Lock()
	defer i.mux.Unlock()
	if time.Since(i.GenerationTime) > i.Timeout {
		log.Println("Generate new UUID")
		i.UUID = uuid.New()
		i.GenerationTime = time.Now()
	}
}

//nolint:nonamedreturns
func (i *Info) GetInfo() (id, hash string, generate time.Time) {
	i.GenerateUUID()
	i.mux.RLock()
	defer i.mux.RUnlock()

	return i.UUID.String(), "not implement", i.GenerationTime
}
