// Get API of auth service put to GraphQL Schema

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AuthService interface {
	Auth(ctx context.Context, username string, password string) (string, error)
}

type authService struct{}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthService() AuthService {
	return &authService{}
}

func loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	URL := os.Getenv("URL")
	AUTH_PORT := os.Getenv("AUTH_PORT")

	return URL, AUTH_PORT
}

func (s *authService) Auth(username string, password string) (string, error) {
	URL, AUTH_PORT := loadConfig()
	authURL := URL + ":" + AUTH_PORT + "/auth"
	
	creds := Credentials{
        Username: username,
        Password: password,
    }

    credsJson, err := json.Marshal(creds)
    if err != nil {
        return "", err
    }

    resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(credsJson))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}