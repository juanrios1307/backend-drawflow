package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"backend/controllers"
)

func Route() *chi.Mux{
	mux := chi.NewMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	})

	//Globals middlewares
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler,
	)

	mux.Get("/getAll", controllers.GetAll)
	mux.Post("/addCode", controllers.GetAll)
	mux.Get("/getById", controllers.GetAll);
	mux.Post("/executeCode", controllers.GetAll)
	return mux

}