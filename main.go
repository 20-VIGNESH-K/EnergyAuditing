package main

import (
	"log"
	"github.com/20-VIGNESH-K/EnergyAuditing/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Fatal(router.Run(":8082"))
}
