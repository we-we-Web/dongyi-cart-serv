package main

import (
	"log"

	"github.com/we-we-Web/dongyi-cart-serv/app/api"
	"github.com/we-we-Web/dongyi-cart-serv/app/infrastructure"
	"github.com/we-we-Web/dongyi-cart-serv/app/repository"
	"github.com/we-we-Web/dongyi-cart-serv/app/usecases"
)

func main() {

	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}

	repo := repository.NewCartRepository(db)
	cartUseCase := usecases.NewCartUseCase(repo)

	router := api.NewRouter(cartUseCase)
	log.Println("Server Start:")
	router.Run(":8080")
}
