package controllers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Selamat datang di Home Page pasar gelap",
	}

	json.NewEncoder(w).Encode(data)
}