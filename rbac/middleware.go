package rbac

import (
	"fmt"

	casbin "github.com/casbin/casbin/v2"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
	echo "github.com/labstack/echo/v4"
)

func Middleware(enforcer *casbin.Enforcer) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
				Enforcer: enforcer,
				UserGetter: func(c echo.Context) (string, error) {
					username, _, _ := c.Request().BasicAuth()
					return username, nil
					// isAdmin, _ := rm.HasLink(username, adminRoleName)
					// if isAdmin {
					// 	return adminRoleName, nil
					// }
					// return webuserRoleName, nil
				},
				ErrorHandler: func(c echo.Context, internal error, proposedStatus int) error {
					err := echo.NewHTTPError(proposedStatus, fmt.Errorf("unauthorized").Error())
					err.Internal = internal
					return err
				},
			})
		}
	}
}
