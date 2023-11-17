package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type MyDB[T any] struct {
	log *log.Logger

	m   sync.Map
	seq atomic.Int64
}

func NewMyDB[T any](logger *log.Logger) *MyDB[T] {
	return &MyDB[T]{
		log: logger,
	}
}

func (d *MyDB[T]) Save(ctx context.Context, data T) (T, error) {
	newID := d.seq.Add(1)

	d.m.Store(newID, data)

	storedData, ok := d.m.Load(newID)
	if !ok {
		var null T
		return null, fmt.Errorf("failed to load data")
	}

	return storedData.(T), nil
}

func (d *MyDB[T]) Query(ctx context.Context) ([]T, error) {
	var res []T

	for i := 1; i < int(d.seq.Load()); i++ {
		gotData, ok := d.m.Load(i)
		if !ok {
			return nil, fmt.Errorf("failed to load data")
		}

		res = append(res, gotData.(T))
	}

	return res, nil
}

func (d *MyDB[T]) Delete(ctx context.Context, id int64) error {
	_, ok := d.m.Load(id)
	if !ok {
		return fmt.Errorf("no data to delete")
	}

	d.m.Delete(id)

	return nil
}
