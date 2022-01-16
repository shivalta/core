package http

import (
	"github.com/labstack/echo/v4"
	"shiva/shiva-auth/factory"
)

func InitHttp() {
	f := factory.InitFactoryHTTP()

	e := echo.New()
	initMiddleware(e)

	v1 := e.Group("api/v1/")
	v1.POST("/login", f.Accounts.Login)

	v1.GET("/users", f.Accounts.GetAll)
	v1.POST("/users", f.Accounts.Create)
	v1.GET("/users/{userId}", f.Accounts.GetById)
	v1.DELETE("/users/{userId}", f.Accounts.Delete)
	v1.PUT("/users/{userId}", f.Accounts.Update)
	err := e.Start(":1111")
	if err != nil {
		return
	}
}