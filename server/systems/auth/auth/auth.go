package auth

import (
	"errors"
	"net/http"

	"training-frontend/package/log"
	"training-frontend/server/systems/auth/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/yudai/pp"
)

const (
	accessTokenCookiePrefix  = "access-token"
	refreshTokenCookiePrefix = "refresh-token"

	//TODO (jwm) get the top secret from configuration file
	jwtSecretKey        = "0ifp5bw(!1j8bq#2bwd24)bn!0$gco6hhoce^!7tmprdaf$1z7" //TODO store this into the configuration
	jwtRefreshSecretKey = "0ifp5bw(!1j7bq#2bwd24)bn!0$gco5hhoce^!7tmprdaf$1z7" //TODO store this into the configuration

	expirationTime = 30 //experation time in minutes
	AuthContextKey = "user"
)

var (
	accessTokenMiddleware echo.MiddlewareFunc
)

// GetJWTSecret returns secreat key
func GetJWTSecret() string {
	return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

func SetTokensAndSetCookies(authResp models.AuthToken, c echo.Context) error {
	accessToken := authResp.AccessToken
	refreshToken := authResp.RefreshToken
	email := authResp.AuthUser
	exp := authResp.ExpireTime
	//accessTokenCookieName := GetCookieKey(accessTokenCookiePrefix, email)
	//accessTokenCookieName := GetCookieKey(accessTokenCookiePrefix, email)
	setTokenCookie(accessTokenCookiePrefix, accessToken, exp, c)
	setUserCookie(email, exp, c)
	//refreshTokenCookieName := GetCookieKey(refreshTokenCookiePrefix, email)
	setTokenCookie(refreshTokenCookiePrefix, refreshToken, exp, c)
	return nil
}

func GetCookieKey(prefix, email string) string {
	return prefix + "-" + email
}

func ClearSession(c echo.Context) {
	_, _, email := GetUserFromContext(c)
	log.Infoln("clearing user session")
	clearUserCookie(c)
	//accessTokenCookieName := GetCookieKey(accessTokenCookiePrefix, email)
	clearTokenCookie(accessTokenCookiePrefix, c)
	refreshTokenCookieName := GetCookieKey(accessTokenCookiePrefix, email)
	clearTokenCookie(refreshTokenCookieName, c)
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func setUserCookie(email string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = email
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func clearTokenCookie(name string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	c.SetCookie(cookie)
}
func clearUserCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func GetUserFromContext(c echo.Context) (userID, campusID int32, emailAddress string) {
	//key := GetCookieKey(accessTokenCookiePrefix, emailAddress)
	accessToken, _ := c.Cookie(accessTokenCookiePrefix)
	if accessToken == nil {
		return 0, 0, ""
	}
	tokenString := accessToken.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //TODO: verify validity
			c.Response().WriteHeader(http.StatusUnauthorized)
			return nil, errors.New("invalid access token")
		}
		return []byte(GetJWTSecret()), nil
	})

	if err != nil {
		pp.Println("error validating token %v", err)
		if err.Error() == "Token is expired" {
			return 0, 0, ""
		}
		return 0, 0, ""
	}
	//var userID int32
	var email string

	if token != nil {
		claims := token.Claims.(jwt.MapClaims)
		userID = int32(claims["id"].(float64))
		campusID = int32(claims["campus_id"].(float64))
		email = claims["email"].(string)
	}
	return userID, campusID, email
}

func GetTokenFromContext(c echo.Context) string {
	accessToken, err := c.Cookie(accessTokenCookiePrefix)
	if accessToken == nil || err != nil {
		return ""
	}
	return accessToken.Value
}
