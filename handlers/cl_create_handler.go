package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/francobottoni/client-api/internal/database"
	"github.com/francobottoni/client-api/models"
)

func (h *CreateClientHandler) SaveClientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

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

func parseRequest(r *http.Request) *models.CreateClientCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateClientCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}
