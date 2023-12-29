package keycloak

import (
	"encoding/json"
	"errors"
	"flowcraft/auth-api/v2/src/config"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type keyCloakError struct {
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

type Keycloak struct {
	clientSecret string
	clientId     string
	realm        string
	resty        *resty.Client
}

func NewKeycloak() *Keycloak {
	configEnv := config.LoadEnv()
	fmt.Println(configEnv)
	return &Keycloak{
		clientSecret: configEnv.ClientSecret,
		clientId:     configEnv.ClientId,
		realm:        configEnv.Realm,
		resty:        resty.New(),
	}
}
func validateHTTPError(res *resty.Response, err error) error {
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		var body_response keyCloakError
		if err = json.Unmarshal(res.Body(), &body_response); err != nil {
			return err
		}
		return errors.New("Unauthorized: " + string(body_response.Error_description))
	}
	return nil
}
func (kc *Keycloak) Login(username string, password string) (response map[string]any, err error) {
	body_req := map[string]string{
		"client_id":     kc.clientId,
		"client_secret": kc.clientSecret,
		"grant_type":    "password",
		"username":      username,
		"password":      password,
		"scope":         "openid",
	}
	res, err := kc.resty.R().SetFormData(body_req).SetHeader("Content-Type", "application/x-www-form-urlencoded").Post(fmt.Sprintf("http://localhost:8085/realms/%s/protocol/openid-connect/token", kc.realm))
	if err = validateHTTPError(res, err); err != nil {
		return nil, err
	}
	var body_response map[string]any
	err = json.Unmarshal(res.Body(), &body_response)
	return body_response, err
}

func (kc *Keycloak) Authenticate(token string) (response map[string]any, err error) {
	res, err := kc.resty.R().SetHeader("Authorization", "Bearer "+token).Get(fmt.Sprintf("http://localhost:8085/realms/%s/protocol/openid-connect/userinfo", kc.realm))
	if err = validateHTTPError(res, err); err != nil {
		fmt.Println(err)
		return nil, err
	}
	var body_response map[string]any
	err = json.Unmarshal(res.Body(), &body_response)
	return body_response, err
}
