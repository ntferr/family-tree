package main

import (
	"family-tree/internal/http/router"
	"family-tree/internal/settings"
	"fmt"
	"log"
)

func main() {
	router := router.SetupRouter()

	serviceAddress := fmt.Sprintf(settings.GetSettings().Host + ":" + settings.GetSettings().Port)

	if err := router.Start(serviceAddress); err != nil {
		log.Fatalln("Error on start router: ", err)
	}
}
