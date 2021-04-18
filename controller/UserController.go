package controller

import (
	"github.com/labstack/echo/v4"
	"golang-echo-rest-api-example/model"
	"log"
	"net/http"
)

func GetAllUser(c echo.Context) error {
	users := []model.User{
		model.User{Name: "Test", Email: "test@gmail.com"},
		model.User{Name: "Test", Email: "test1@gmail.com"},
	}
	return c.JSON(http.StatusOK, users)
}

func SaveUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	log.Printf("User Id : %s", id)

	user := model.User{
		Name:  "Test",
		Email: "test@gmail.com",
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	log.Printf("Updated User Id : %s", id)

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	log.Printf("Deleted User Id : %s", id)

	return c.NoContent(http.StatusNoContent)
}
