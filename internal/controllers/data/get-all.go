package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorvk/starterapp/internal/initializers"
	models "github.com/gorvk/starterapp/internal/models/data"
	"github.com/gorvk/starterapp/internal/types"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	auth := initializers.GetAuthInstance()
	userClaim, err := auth.VerifyIDToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	rows, err := models.GetAll(userClaim)
	if err != nil {
		fmt.Println("response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := []types.Data{}
	defer rows.Close()
	for rows.Next() {
		row := types.Data{}
		rows.Scan(
			&row.Id,
			&row.UserId,
		)
		data = append(data, row)
	}

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
