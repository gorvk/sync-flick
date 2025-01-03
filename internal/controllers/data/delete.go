package controllers

import (
	"net/http"

	"github.com/gorvk/starterapp/internal/initializers"
	models "github.com/gorvk/starterapp/internal/models/data"
	"github.com/gorvk/starterapp/internal/utils"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	auth := initializers.GetAuthInstance()
	userClaim, err := auth.VerifyIDToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	id := r.PathValue("id")
	err = models.Delete(id, userClaim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := utils.ConstructResponse(true, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
