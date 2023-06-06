package handler

import (
	"api/model"
	"encoding/json"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {
	listProducts := model.GetAllProducts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listProducts)
}
