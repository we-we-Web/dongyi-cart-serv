package main

import (
	"log"

	"github.com/we-we-Web/dongyi-cart-serv/app/api"
	"github.com/we-we-Web/dongyi-cart-serv/app/infrastructure"
)

func main() {

	_, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}

	// repo := repository.NewChatRepository(db)
	// chatUseCase := usecase.NewMsgUseCase(repo)

	router := api.NewRouter()
	log.Println("Server Start:")
	router.Run(":8080")
}
