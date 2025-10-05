package query

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type ValidateAccessTokenRequest struct {
	Token string `json:"token"`
}
type ValidateAccessTokenResponse struct {
	Message string `json:"message"`
}
type ValidateAccessTokenHandler struct {
}

func NewValidateAccessTokenHandler() *ValidateAccessTokenHandler {
	return nil
}

func (h *ValidateAccessTokenHandler) Handle(ctx context.Context, req *ValidateAccessTokenRequest) (*ValidateAccessTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenSignatureInvalid
		}
		return []byte(viper.GetString("server.secret_key")), nil

	})

	//burda hata giriyor çünkü exp float64 olması gerekirken ben time.Time gönderiyorum
	if err != nil {
		return nil, ErrTokenNotProvided
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenClaimsInvalid
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, ErrTokenExpInvalid
	}

	if int64(expFloat)-time.Now().Unix() < 0 {
		return nil, ErrTokenExpired
	}

	nick, ok := claims["nick"].(string)
	if !ok {
		return nil, ErrTokenNickNotFound
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil, ErrTokenRoleNotFound
	}
	return &ValidateAccessTokenResponse{
		Message: fmt.Sprintf("Access granted as %s and the role is %s", nick, role),
	}, nil
}
