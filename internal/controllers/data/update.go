package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/starterapp/internal/initializers"
	models "github.com/gorvk/starterapp/internal/models/data"
	"github.com/gorvk/starterapp/internal/types"
	"github.com/gorvk/starterapp/internal/utils"
)

func Update(w http.ResponseWriter, r *http.Request) {
	auth := initializers.GetAuthInstance()
	userClaim, err := auth.VerifyIDToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var input = types.Data{}
	err = json.Unmarshal(d, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = models.Update(input, userClaim)
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
