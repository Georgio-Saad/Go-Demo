package response

import "todogorest/models"

type AuthResponse struct {
	models.User
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}
