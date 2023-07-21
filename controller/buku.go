package controller

import (
	"len_recruitment/database"
	"len_recruitment/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBuku(c echo.Context) error {
	buku := model.Buku{}

	c.Bind(&buku)

	if err := database.DB.Save(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create buku",
		"status":  201,
		"data":    buku,
	})
}

func GetAllBuku(c echo.Context) error {
	buku := []model.Buku{}

	if err := database.DB.Find(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all buku",
		"status":  200,
		"data":    buku,
	})
}

func GetBukuByID(c echo.Context) error {
	buku := model.Buku{}

	bukuID := c.Param("id")

	if err := database.DB.First(&buku, bukuID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get buku by id",
		"status":  200,
		"data":    buku,
	})
}

func UpdateBukuByID(c echo.Context) error {
	buku := model.Buku{}

	bukuID := c.Param("id")

	if err := database.DB.First(&buku, bukuID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&buku)

	if err := database.DB.Save(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update buku by id",
		"status":  200,
		"data":    buku,
	})
}

func DeleteBukuByID(c echo.Context) error {
	buku := model.Buku{}

	bukuID := c.Param("id")

	if err := database.DB.Delete(&buku, bukuID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update buku by id",
		"status":  200,
		"data":    buku,
	})
}
