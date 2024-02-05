package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevlar-ashu/simple-learning/pkg/config"
	"github.com/kevlar-ashu/simple-learning/pkg/handlers"
	"github.com/kevlar-ashu/simple-learning/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting the server on port", portNumber)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf(fmt.Sprintf("Couldn't start the Server on port %s\n", portNumber))
	}

}
