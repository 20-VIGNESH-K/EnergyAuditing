package main

import (
	"log"
	"os"

	"github.com/20-VIGNESH-K/EnergyAuditing/routes"
)

func main() {
	router := routes.SetupRouter()
	port := os.Getenv("port")
	if port == "" {
        port = "8080"
    }
	log.Fatal(router.Run(":"+port))
}
