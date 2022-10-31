package main

import (
	"fmt"
	"log"

	"github.com/checkaayush/authware/handler"
	"github.com/checkaayush/authware/repository/inmem"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Seed data:
// Users: U1
// Metrics: M1, M2
// Blocks: B1 (M1), B2 (M2)
// Apps: A1 (B1, B2)

func hasAccessToMetric(enforcer *casbin.Enforcer, user, metric string) (bool, error) {
	metricReaderRole := fmt.Sprintf("%s_Readers", metric)
	return enforcer.HasRoleForUser(user, metricReaderRole)
}

func showBlock(enforcer *casbin.Enforcer, user, block string) (bool, error) {
	blockReaderRole := fmt.Sprintf("%s_Readers", block)
	return enforcer.HasRoleForUser(user, blockReaderRole)
}

func hasAccessToBlock(enforcer *casbin.Enforcer, user, block string, metric string) (bool, error) {
	metricOk, err := hasAccessToMetric(enforcer, user, metric)
	if err != nil || !metricOk {
		return false, fmt.Errorf("you don't have access to the associated metric")
	}

	return true, nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	e := echo.New()

	// Add middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	enforcer, err := casbin.NewEnforcer("auth_model.conf", "auth_policy.csv")
	if err != nil {
		e.Logger.Fatal(err)
	}

	user1, metric1, block1 := "U1", "M1", "B1"
	user2, _, _ := "U2", "M2", "B2"
	ok, err := hasAccessToMetric(enforcer, user1, metric1)
	handleErr(err)
	log.Printf("%s allowed %s: %t", user1, metric1, ok)

	ok, err = hasAccessToMetric(enforcer, user2, metric1)
	handleErr(err)
	log.Printf("%s allowed %s: %t", user2, metric1, ok)

	ok, err = showBlock(enforcer, user1, block1)
	handleErr(err)
	log.Printf("%s should see %s: %t", user1, block1, ok)

	ok, err = hasAccessToBlock(enforcer, user1, block1, metric1)
	handleErr(err)
	log.Printf("%s should see visualisation for %s: %t", user1, block1, ok)

	ok, err = showBlock(enforcer, user2, block1)
	handleErr(err)
	log.Printf("%s should see %s: %t", user2, block1, ok)

	// rm := oktarolemanager.NewRoleManager("dev-17237792",
	// 	"00xtDNKekzJi9R6kGxsNn5sBk-QroWkwk4USstaT9w", true)
	// enforcer.SetRoleManager(rm)
	// enforcer.LoadPolicy()
	// e.Use(casbin_mw.Middleware(enforcer))

	// Initialize handler
	h := handler.NewHandler(inmem.NewInMemRepository())

	// roles, err := rm.GetRoles("checkaayush@gmail.com")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// log.Println(roles)
	// ok, err := enforcer.Enforce("checkaayush@gmail.com", "/api/v1/admin", "GET")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(ok)
	// users := createUsers()

	// users, err := rm.GetUsers("checkaayush@gmail.com")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// log.Println(users)
	// log.Println("getting users with role Admin")
	// roles, err := rm.GetRoles("checkaayush@gmail.com")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// log.Println(roles)
	// If our role manager relies on Casbin policy (like reading "g"
	// policy rules), then we have to set the role manager before loading
	// policy.
	//
	// Otherwise, we can set the role manager at any time, because role
	// manager has nothing to do with the adapter.
	// enforcer.LoadPolicy()

	// Check the permission.
	// Casbin's subject (user) name uses the Okta user's login field (aka Email address).
	// Casbin's role name uses the Okta group's name field (like "Admin", "Everyone").
	// e.Enforce("alice@test.com", "data1", "read")

	// Routes
	v1 := e.Group("/api/v1")

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
	v1.DELETE("/blocks/:id", h.DeleteBlock)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
