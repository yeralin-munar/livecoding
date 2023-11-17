package main

import (
	"context"
	"fmt"
	"livecoding/internal/usecase"
	"log"
)

/*
Сервис для конструктора кошек
- Смешивать разные породы -
- Породы кошек и их связь с генами
- Есть API провайдер:
	- О породах
	- О генах

		Пример:
	Cat breads:

		[{ name: “Abyssinian”, primaryGen: “z-kar” }, {name: “Balinese”, primaryGen: “k-pex”}, { name: “Devon Rex”, primaryGen: “t-rex” }]

	Gens:

		[{ name: “z-kar”, factor: 3}, {name: “k-pex”, factor: 10 }, { name: “t-rex”, factor: 15}]

Задача:
- Сымитировать вызов к API провайдеру
- Сохранить в БД

- Реализовать 2 API:
  - Получение списка всех генов и связанных с ними породами кошек
	- Удаление гена по ключу (из БД). Если в каких-то породах он является доминирующим, то эту породу удалить


Тесты:
1)
- Получить список генов
- Удалить ген z-kar
- Получить список


*/

func main() {
	fmt.Println("Hello, World!")

	logger := log.Logger{}
	catUsecase := usecase.NewCatUsecase(&logger)

	res, err := catUsecase.GetGensAndBreads(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
