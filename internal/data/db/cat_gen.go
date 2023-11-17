package db

import (
	"context"
	"fmt"
	"log"

	"livecoding/internal/usecase/dao"
)

type CatGenStore struct {
	log *log.Logger
	db  *MyDB[*dao.Gen]
}

func NewCatGenStore(logger *log.Logger) *CatGenStore {
	db := NewMyDB[*dao.Gen](logger)

	return &CatGenStore{
		log: logger,
		db:  db,
	}
}

func (c *CatGenStore) Query(ctx context.Context) ([]*dao.Gen, error) {
	res, err := c.db.Query(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query cat gens: %w", err)
	}

	c.log.Printf("Query cat gens: %v", res)

	return res, nil
}

func (c *CatGenStore) Save(ctx context.Context, data []*dao.Gen) ([]*dao.Gen, error) {
	var res []*dao.Gen

	for i := range data {
		item := data[i]

		storedItem, err := c.db.Save(ctx, item)
		if err != nil {
			return nil, fmt.Errorf("failed to save cat gen %v: %w", *item, err)
		}

		res = append(res, storedItem)
	}

	c.log.Printf("Save cat gens: %v", res)

	return res, nil
}
