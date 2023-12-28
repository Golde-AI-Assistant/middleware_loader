// Get API of auth service put to GraphQL Schema

package service_api

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

type AuthService interface {
	Auth(ctx context.Context, username string, password string) (string, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (a *authService) Auth(ctx context.Context, username string, password string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	log.Printf("Username: %s", username)
	return username, nil
	// authUrl := fmt.Sprintf("%s/auth", os.Getenv("AUTH_SERVICE_URL"))

	// data := map[string]string{
	// 	"username": username,
	// 	"password": password,
	// }

	// payload, err := json.Marshal(data)
	// if err != nil {
	// 	return "", err
	// }

	// req, err := http.NewRequest("POST", authUrl, bytes.NewBuffer(payload))
	// if err != nil {
	// 	return "", err
	// }

	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return "", err
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }

	// var result map[string]interface{}
	// json.Unmarshal([]byte(body), &result)

	// return result["token"].(string), nil
}
