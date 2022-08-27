package main


import (
	"net/http"

	"backend/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	
)

func main() {

	r :=chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	})

	r.Use(cors.Handler)

	r.Route("/" , func(r chi.Router){
		r.Get("/" , controllers.GetAll)
		r.Post("/" , controllers.Add)
		r.Get("/exec",controllers.Execute)
		r.Route("/{id}", func(r chi.Router){
			r.Get("/", controllers.GetOne)
		})
	})

	println("Listen on 9000")
	http.ListenAndServe(":9000", r)
}

