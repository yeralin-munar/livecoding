package cat_provider

import "log"

type CatBread struct {
	Name       string
	PrimaryGen string
}

type Gen struct {
	Name   string
	Factor int
}

type CatProvider struct {
	log *log.Logger
}

func NewCatProvider(logger log.Logger) *CatProvider {
	return &CatProvider{
		log: &logger,
	}
}

func (c *CatProvider) GetCatBreads() ([]*CatBread, error) {
	// [{ name: “Abyssinian”, primaryGen: “z-kar” }, {name: “Balinese”, primaryGen: “k-pex”}, { name: “Devon Rex”, primaryGen: “t-rex” }]

	data := []*CatBread{
		{
			Name:       "Abyssinian",
			PrimaryGen: "z-kar",
		},
		{
			Name:       "Balinese",
			PrimaryGen: "k-pex",
		},
		{
			Name:       "Devon Rex",
			PrimaryGen: "t-rex",
		},
	}

	c.log.Println("Got cat breads: ", data)

	return data, nil
}

// [{ name: “z-kar”, factor: 3}, {name: “k-pex”, factor: 10 }, { name: “t-rex”, factor: 15}]
func (c *CatProvider) GetGens() ([]*Gen, error) {
	data := []*Gen{
		{
			Name:   "z-kar",
			Factor: 3,
		},
		{
			Name:   "k-pex",
			Factor: 10,
		},
		{
			Name:   "t-rex",
			Factor: 15,
		},
	}

	c.log.Println("Got gens: ", data)

	return data, nil
}
