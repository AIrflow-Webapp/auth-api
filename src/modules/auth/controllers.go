package auth

type AuthController struct {
	service *AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		service: NewAuthService(),
	}
}

func (ac *AuthController) Login(loginDTO LoginDTO) (int, map[string]any) {
	jwt, error := ac.service.login(loginDTO.Username, loginDTO.Password)
	if error != nil {
		return 500, map[string]any{
			"error": error.Error(),
		}
	}
	return 200, jwt
}

func (ac *AuthController) Authenticate(token string) (int, map[string]any) {
	jwt, error := ac.service.authenticate(token)
	if error != nil {
		return 500, map[string]any{
			"error": error.Error(),
		}
	}
	return 200, jwt
}
