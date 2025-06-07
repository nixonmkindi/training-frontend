package auth

import (
	"net/http"
	"training-frontend/package/log"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthJWT() echo.MiddlewareFunc {
	if accessTokenMiddleware != nil {
		return accessTokenMiddleware
	}
	accessTokenMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &JWTCustomClaims{},
		SigningKey:  []byte(GetJWTSecret()),
		TokenLookup: "cookie:" + accessTokenCookiePrefix,
		//TokenLookup:             "header:" + echo.HeaderAuthorization,
		ErrorHandlerWithContext: JWTErrorChecker,
		Skipper:                 SkipperLoginCheck,
		AuthScheme:              "Bearer",
	})
	return accessTokenMiddleware
}

// SkipperLoginCheck register all routes that do not need login
func SkipperLoginCheck(c echo.Context) bool {
	if strings.HasSuffix(c.Path(), "/auth/login") ||
		strings.HasSuffix(c.Path(), "/auth/register") ||
		strings.HasSuffix(c.Path(), "/auth/register/create") ||
		strings.HasSuffix(c.Path(), "/auth/users/select-campus") ||
		strings.HasSuffix(c.Path(), "/auth/users/update-campus") ||
		strings.HasPrefix(c.Path(), "/images") ||
		strings.HasPrefix(c.Path(), "/favicon.ico") ||
		strings.HasPrefix(c.Path(), "/uploads") ||
		strings.HasPrefix(c.Path(), "/error") ||
		strings.Contains(c.Path(), "adminlte") {
		return true
	}
	return false
}

// JWTErrorChecker will be executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	log.Errorf("not allowed to access this url: %v, err: %v\n", c.Path(), err)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
