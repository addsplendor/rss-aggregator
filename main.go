package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hello, world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:", portString)

	router := chi.NewRouter()

	srv:= &http.Server{
		Handler: router,
		Addr: ":"+ portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}
