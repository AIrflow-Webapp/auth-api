package auth

import (
	"flowcraft/auth-api/v2/src/connectors/keycloak"
	"fmt"
)

type AuthService struct {
	kc *keycloak.Keycloak
}

func NewAuthService() *AuthService {
	return &AuthService{
		kc: keycloak.NewKeycloak(),
	}
}

func (as *AuthService) login(username string, password string) (jwt map[string]any, err error) {
	jwt, err = as.kc.Login(username, password)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

func (as *AuthService) authenticate(token string) (jwt map[string]any, err error) {
	jwt, err = as.kc.Authenticate(token)
	fmt.Println(jwt, err)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}
