package main

import (
	"coffe-shop/database"
	"coffe-shop/services"
	auth "coffe-shop/services/authentication"
	"coffe-shop/services/guard"
	places2 "coffe-shop/services/places"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.LoadDatabase()
	e := echo.New()
	cfg := services.GetConfig()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// auth
	e.POST("/login", auth.LoginUser)
	e.POST("/register", auth.RegisterUser)

	// test
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	places := e.Group("/place")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(services.CustomClaims)
		},
		SigningKey: []byte(cfg.JwtSecret),
	}
	places.GET("/:id", places2.GetPlaceByID)
	places.GET("", places2.GetPlaces)
	places.POST("", guard.OwnerGuard(places2.InsertPlace), echojwt.WithConfig(config))
	places.PUT("", guard.AdminGuard(places2.UpdatePlace), echojwt.WithConfig(config))
	places.DELETE("", guard.OwnerGuard(places2.DeletePlace), echojwt.WithConfig(config))

	// comment
	places.GET("/comment/:id", places2.GetCommentByPlaceID)
	places.POST("/comment", guard.AuthGuard(places2.AddComment), echojwt.WithConfig(config))

	e.Logger.Fatal(e.Start(":1323"))
}

