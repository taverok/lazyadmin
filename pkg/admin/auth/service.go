package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/taverok/lazyadmin/pkg/admin/auth/provider"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/jwtutils"
)

type Service struct {
	Config   *config.Config
	Provider provider.Provider
}

func (it *Service) AuthJwt(request LoginRequest) (string, error) {
	principal, err := it.Provider.Authenticate(request.User, request.Pass)
	if err != nil {
		return "", err
	}

	claims := jwtutils.JwtClaims{
		Uuid:             principal.Name,
		Role:             principal.Role,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	return claims.GenerateJwtToken(it.Config.Auth.JwtSecret)
}
