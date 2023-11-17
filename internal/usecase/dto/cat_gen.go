package dto

import "livecoding/internal/usecase/dao"

type CatGen struct {
	Gen    *dao.Gen
	Breads []*dao.CatBread
}
