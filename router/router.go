package router

import (
	"net/http"

	"github.com/VolodymyrVoronov/cats-back/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/api/cats", controllers.GetAllCats)
	router.Get("/api/cats/{id}", controllers.GetCatByID)
	router.Post("/api/cat", controllers.CreateCat)
	router.Delete("/api/cat/{id}", controllers.DeleteCatByID)

	return router
}
