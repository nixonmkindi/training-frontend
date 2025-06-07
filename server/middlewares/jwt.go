package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/services/usecase/audit_trails"
	"training-frontend/server/systems/auth/auth"
	"training-frontend/server/systems/helpers"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// JWT Middleware
func JWT() echo.MiddlewareFunc {
	return auth.AuthJWT()
}

// CheckLogin middleware
func CheckAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if auth.SkipperLoginCheck(c) {
				return next(c)
			}
			u := c.Get("user").(*jwt.Token)
			if u == nil {
				return c.String(http.StatusUnauthorized, "unable to get token")
			}
			if !u.Valid {
				return c.Redirect(http.StatusFound, "/auth/login")
			}
			claims := u.Claims.(*auth.JWTCustomClaims)
			if claims == nil {
				return c.Redirect(http.StatusFound, "/auth/login")
			}

			//added as part of audit trail
			url := c.Request().URL
			buf := bytes.Buffer{}
			tee := io.TeeReader(c.Request().Body, &buf)
			body, _ := io.ReadAll(tee)
			c.Request().Body = io.NopCloser(&buf)
			auditData := AuditTrails{
				UserID:    claims.ID,
				Action:    c.Request().Method,
				Url:       url.Path,
				IPAddress: url.Host + "-" + c.RealIP(),
				Client:    c.Request().UserAgent(),
				Data:      string(body),
			}
			auditService := audit_trails.NewService()
			auditService.CreateAuditTrails(auditData.UserID, auditData.IPAddress, auditData.Client, auditData.Action, auditData.Url, auditData.Data)
			// end of audit trail

			p, err := helpers.HasPermission(claims.Email, url.String())
			if p && err == nil {
				log.Infoln("Access Granted")
				return next(c)
			} else {
				helpers.ClearCache(claims.Email)
				return c.Redirect(http.StatusFound, "/auth/login")
			}
		}
	}
}
