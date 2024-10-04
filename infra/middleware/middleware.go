package middleware

import (
	"base-api/constants"
	"base-api/utils"
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

type TokenMiddlewareInterface interface {
	TokenAuthorize(handlerFunc http.Handler) http.Handler
}

type tokenMiddlewareObj struct {
	JWTService JWTInterface
}

func NewTokenMiddleware(jwt JWTInterface) TokenMiddlewareInterface {
	return &tokenMiddlewareObj{
		jwt,
	}
}

func (m *tokenMiddlewareObj) TokenAuthorize(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get(constants.Authorization)
		if tokenHeader == "" {
			err := constants.ErrTokenIsRequired
			log.Error(err)
			res := utils.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Err:  utils.STATUS_UNAUTHORIZED,
				Msg:  err.Error(),
			}
			res.JSONResponse(w)
			return
		}

		claims, err := m.JWTService.ExtractJWTClaims(r.Context(), tokenHeader)
		if err != nil {
			code := http.StatusInternalServerError
			errMsg := utils.STATUS_INTERNAL_ERR
			if err == constants.ErrTokenAlreadyExpired || err == constants.ErrTokenReplaced || err == constants.ErrTokenInvalid {
				code = http.StatusUnauthorized
				errMsg = utils.STATUS_UNAUTHORIZED
			}

			log.Error(err)
			res := utils.SetResponseJSON(code, nil, errMsg, err.Error())
			res.JSONResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), tokenCtxKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

// GetTokenFromContext return Token Object inside context data
func GetTokenFromContext(ctx context.Context) *JWTClaims {
	raw, _ := ctx.Value(tokenCtxKey).(*JWTClaims)
	return raw
}
