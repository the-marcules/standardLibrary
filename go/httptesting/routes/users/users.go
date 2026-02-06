package users

import (
	"encoding/json"
	"net/http"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userfile := os.Getenv("USER_FILE")

	userFileContents, err := os.ReadFile(userfile)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users Users
	err = json.Unmarshal(userFileContents, &users)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(userFileContents)
}
