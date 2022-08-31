package main


import (
	"net/http"

	"backend/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	
)

func main() {

	//Declaración del router en libreria chi
	r :=chi.NewRouter()

	//Habilita opciones del CORS para permitir conexión con browser 
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	})

	//Router utiliza las opciones del cors
	r.Use(cors.Handler)

	//Declaración de rutas para ejecutar endpoints
	r.Route("/" , func(r chi.Router){
		r.Get("/" , controllers.GetAll)
		r.Post("/" , controllers.Add)
		r.Get("/exec",controllers.Execute)
		r.Route("/{id}", func(r chi.Router){
			r.Get("/", controllers.GetOne)
		})
	})

	//Escucha y abre servidor en puerto 9000
	println("Listen on 9000")
	http.ListenAndServe(":9000", r)
}

