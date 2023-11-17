package db

import (
	"context"
	"fmt"
	"livecoding/internal/usecase/dao"
	"log"
)

type CatBreadStore struct {
	log *log.Logger
	db  *MyDB[*dao.CatBread]
}

func NewCatBreadStore(logger *log.Logger) *CatBreadStore {
	db := NewMyDB[*dao.CatBread](logger)

	return &CatBreadStore{
		log: logger,
		db:  db,
	}
}

func (c *CatBreadStore) Query(ctx context.Context) ([]*dao.CatBread, error) {
	res, err := c.db.Query(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query cat breads: %w", err)
	}

	c.log.Printf("Query cat breads: %v", res)

	return res, nil
}

func (c *CatBreadStore) Save(ctx context.Context, data []*dao.CatBread) ([]*dao.CatBread, error) {
	var res []*dao.CatBread

	for i := range data {
		item := data[i]

		storedItem, err := c.db.Save(ctx, item)
		if err != nil {
			return nil, fmt.Errorf("failed to save cat bread %v: %w", *item, err)
		}

		res = append(res, storedItem)
	}

	c.log.Printf("Save cat breads: %v", res)

	return res, nil
}
