package dao

import (
	"livecoding/internal/data/cat_provider"
)

type CatBread struct {
	Name       string
	PrimaryGen string
}

type Gen struct {
	Name   string
	Factor int
}

func ConvertCPCatBreadsToCatBreads(in []*cat_provider.CatBread) []*CatBread {
	res := make([]*CatBread, len(in))

	for i := range in {
		newCatBread := &CatBread{
			Name:       in[i].Name,
			PrimaryGen: in[i].PrimaryGen,
		}

		res[i] = newCatBread
	}

	return res
}

func ConvertCPGensToGens(in []*cat_provider.Gen) []*Gen {
	res := make([]*Gen, len(in))

	for i := range in {
		newGen := &Gen{
			Name:   in[i].Name,
			Factor: in[i].Factor,
		}

		res[i] = newGen
	}

	return res
}
