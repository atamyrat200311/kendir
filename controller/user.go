package controller

import (
	"encoding/json"
	"net/http"

	"kendir-mini/db"
	"kendir-mini/dto"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input dto.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if len(input.Name) < 5 {
		http.Error(w, "Name must be 5 letters at least", http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(
		"INSERT INTO users(name) VALUES($1) RETURNING id",
		input.Name,
	).Scan(&input.Id)

	if err != nil {
		http.Error(w, "Insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []dto.User
	for rows.Next() {
		var u dto.User
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}
