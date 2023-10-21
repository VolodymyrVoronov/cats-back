package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/VolodymyrVoronov/cats-back/database"
	"github.com/VolodymyrVoronov/cats-back/router"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	port := app.Config.Port

	fmt.Printf("Listening on port %d\n", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portInt, _ := strconv.Atoi(port)

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Cannot connect to database %s", err)
	}

	defer db.Client.Disconnect()

	config := Config{
		Port: portInt,
	}

	app := &Application{
		Config: config,
	}

	err = app.Serve()
	if err != nil {
		log.Fatalf("Cannot start server %s", err)
	}
}
