package controllers

import (
	"fmt"
	"mercari-clone-backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ItemIndex(c echo.Context) error {
	db := models.Connect()
	defer db.Close()

  var items []models.Item
	db.Find(&items)
	return c.JSON(http.StatusOK, items)
}

func ItemShow(c echo.Context) error {
	db := models.Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	item := new(models.Item)
	db.First(&item, id)

	return c.JSON(http.StatusOK, item)
}

func ItemNew(c echo.Context) error {
	db := models.Connect()
	defer db.Close()

	newItem := new(models.Item)
	if err := c.Bind(newItem); err != nil {
		return err
	}
  fmt.Printf("%v", newItem)
	db.Create(&newItem)

	return c.JSON(http.StatusOK, newItem)
}

func ItemEdit(c echo.Context) error {
	db := models.Connect()
	defer db.Close()

	var item models.Item
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	db.First(&item, id)

	return c.NoContent(http.StatusOK)
}

func ItemDelete(c echo.Context) error {
	db := models.Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	item := new(models.Item)
	db.Delete(&item, id)
	return c.NoContent(http.StatusOK)
}
