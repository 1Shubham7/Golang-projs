package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {

	godotenv.Load(".env") //by defalut go imports the .env file
	

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviornment")
	}
	fmt.Println("Port: ", portString)
	
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	srv := &http.Server{
		Handler : router,
		Addr : ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}