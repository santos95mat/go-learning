package httpRequest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Number    string    `json:"number"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUsers() {
	res, err := http.Get("http://localhost:3000/v1/user")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var users []User
	err = json.Unmarshal(body, &users)

	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		log.Printf("%+v\n", user)
	}
}
