package routes

import (
	"net/http"

	controllers "github.com/gorvk/starterapp/internal/controllers/data"
)

func init() {
	http.HandleFunc("GET /api/data", controllers.GetAll)
	http.HandleFunc("GET /api/data/{id}", controllers.GetById)
	http.HandleFunc("POST /api/data", controllers.Create)
	http.HandleFunc("PUT /api/data", controllers.Update)
	http.HandleFunc("DELETE /api/data/{id}", controllers.Delete)
}
