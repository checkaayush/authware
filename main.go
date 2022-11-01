package main

import (
	"fmt"

	"github.com/checkaayush/authware/handler"
	"github.com/checkaayush/authware/rbac"
	"github.com/checkaayush/authware/repository/inmem"

	casbin "github.com/casbin/casbin/v2"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
	echo "github.com/labstack/echo/v4"
	echo_mw "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Add basic middlewares
	e.Use(echo_mw.Logger())
	e.Use(echo_mw.Recover())

	enforcer, err := casbin.NewEnforcer("auth_model.conf", "auth_policy.csv")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Add Authorization middleware
	e.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
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
	}))

	// Initialize storage
	repo := inmem.NewInMemRepository()

	// Initialize RBAC
	auth, err := rbac.New(enforcer)
	if err != nil {
		e.Logger.Fatalf("failed to initialize rbac: %s", err)
	}

	// Initialize handler
	h := handler.NewHandler(repo, auth)

	// Routes
	v1 := e.Group("/api/v1")
	v1.GET("/health", h.Health)

	v1.POST("/users", h.InviteUser)
	v1.GET("/users", h.ListUsers)
	v1.DELETE("/users/:id", h.DeleteUser)

	v1.POST("/metrics", h.CreateMetric)
	v1.GET("/metrics", h.ListMetrics)
	v1.DELETE("/metrics/:id", h.DeleteMetric)

	v1.POST("/apps", h.CreateApp)
	v1.GET("/apps", h.ListApps)
	v1.DELETE("/apps/:id", h.DeleteApp)

	v1.POST("/blocks", h.CreateBlock)
	v1.GET("/blocks", h.ListBlocks)
	v1.GET("/blocks/:id", h.GetBlockByID)
	v1.DELETE("/blocks/:id", h.DeleteBlock)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
