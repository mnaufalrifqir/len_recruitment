package controller

import (
	"len_recruitment/database"
	"len_recruitment/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMahasiswa(c echo.Context) error {
	mahasiswa := model.Mahasiswa{}

	c.Bind(&mahasiswa)

	if err := database.DB.Save(&mahasiswa).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create mahasiswa",
		"status":  201,
		"data":    mahasiswa,
	})
}

func GetAllMahasiswa(c echo.Context) error {
	mahasiswa := []model.Mahasiswa{}

	if err := database.DB.Find(&mahasiswa).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all mahasiswa",
		"status":  200,
		"data":    mahasiswa,
	})
}

func GetMahasiswaByID(c echo.Context) error {
	mahasiswa := model.Mahasiswa{}

	mahasiswaID := c.Param("id")

	if err := database.DB.First(&mahasiswa, mahasiswaID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get mahasiswa by id",
		"status":  200,
		"data":    mahasiswa,
	})
}

func UpdateMahasiswaByID(c echo.Context) error {
	mahasiswa := model.Mahasiswa{}

	mahasiswaID := c.Param("id")

	if err := database.DB.First(&mahasiswa, mahasiswaID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&mahasiswa)

	if err := database.DB.Save(&mahasiswa).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update mahasiswa by id",
		"status":  200,
		"data":    mahasiswa,
	})
}

func DeleteMahasiswaByID(c echo.Context) error {
	mahasiswa := model.Mahasiswa{}

	mahasiswaID := c.Param("id")

	if err := database.DB.Delete(&mahasiswa, mahasiswaID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update mahasiswa by id",
		"status":  200,
		"data":    mahasiswa,
	})
}
