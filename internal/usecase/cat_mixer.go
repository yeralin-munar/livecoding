package usecase

import (
	"context"
	"livecoding/internal/data/db"
	"livecoding/internal/usecase/dao"
	"livecoding/internal/usecase/dto"
	"log"
)

type CatUsecase struct {
	log           *log.Logger
	catBreadStore *db.CatBreadStore
	catGenStore   *db.CatGenStore
}

func NewCatUsecase(logger *log.Logger) *CatUsecase {
	catBreadStore := db.NewCatBreadStore(logger)
	catGenStore := db.NewCatGenStore(logger)

	return &CatUsecase{
		log:           logger,
		catBreadStore: catBreadStore,
		catGenStore:   catGenStore,
	}
}

// Получение списка всех генов и связанных с ними породами кошек
func (c *CatUsecase) GetGensAndBreads(ctx context.Context) ([]*dto.CatGen, error) {
	gens, err := c.catGenStore.Query(ctx)
	if err != nil {
		return nil, err
	}

	breads, err := c.catBreadStore.Query(ctx)
	if err != nil {
		return nil, err
	}

	var res []*dto.CatGen
	for i := range gens {
		gen := gens[i]

		var breadsDto []*dao.CatBread
		for j := range breads {
			bread := breads[j]

			if bread.PrimaryGen == gen.Name {
				breadsDto = append(breadsDto, bread)
			}
		}

		newCatGen := &dto.CatGen{
			Gen:    gen,
			Breads: breadsDto,
		}

		res = append(res, newCatGen)
	}

	return res, nil
}

// Удаление гена по ключу (из БД). Если в каких-то породах он является доминирующим, то эту породу удалить

// func (c *CatUsecase) DeleteGen(ctx context.Context, genName string) error {
// 	gens, err := c.catGenStore.Query(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	// Удалить ген

// 	for i := range gens {
// 		gen := gens[i]

// 		if gen.Name == genName {
// 			// Удаляем
// 			break
// 		}
// 	}
// }
