package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/francobottoni/client-api/internal/database"
	"github.com/francobottoni/client-api/models"
)

func (h *CreateClientHandler) SaveClientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	res, err := h.Create(&models.CreateClientCMD{
		DNI:           39098710,
		Name:          "franco",
		LastName:      "bottoni",
		CountryOrigin: "argentina",
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create client"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)

}

type CreateClientHandler struct {
	models.CreateClientI
}

func NewCreateClientHandler(c *database.MySqlClient) *CreateClientHandler {
	return &CreateClientHandler{models.NewClientCreate(c)}
}
