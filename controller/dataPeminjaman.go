package controller

import (
	"len_recruitment/database"
	"len_recruitment/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateDataPeminjaman(c echo.Context) error {
	type Request struct {
		MahasiswaID  uint      `json:"mahasiswa_id"`
		BukuID       uint      `json:"buku_id"`
		TanggalBatas time.Time `json:"tanggal_batas"`
	}
	dataPeminjaman := Request{}
	mahasiswa := model.Mahasiswa{}
	buku := model.Buku{}
	peminjamanMahasiswa := model.DataPeminjaman{}

	c.Bind(&dataPeminjaman)

	if err := database.DB.First(&mahasiswa, dataPeminjaman.MahasiswaID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.First(&buku, dataPeminjaman.BukuID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Order("tanggal_peminjaman desc").First(&peminjamanMahasiswa).Where("mahasiswa_id = ?", dataPeminjaman.MahasiswaID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if peminjamanMahasiswa.TanggalPeminjaman.Month()+1 == time.Now().Month() {
		mahasiswa.TotalPinjaman = 0
	}

	if mahasiswa.TotalPinjaman < 10 && buku.Kuantitas > 0 {
		if err := database.DB.Model(&model.DataPeminjaman{}).Create([]map[string]interface{}{
			{"created_at": time.Now(),
				"updated_at":    time.Now(),
				"mahasiswa_id":  dataPeminjaman.MahasiswaID,
				"buku_id":       dataPeminjaman.BukuID,
				"tanggal_batas": dataPeminjaman.TanggalBatas},
		}).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		mahasiswa.TotalPinjaman++
		buku.Kuantitas--
		database.DB.Save(&buku)
		database.DB.Save(&mahasiswa)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create data peminjaman",
			"status":  201,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "failed create data peminjaman",
			"status":  400,
		})
	}
}

func PengembalianBuku(c echo.Context) error {
	dataPeminjaman := model.DataPeminjaman{}
	mahasiswa := model.Mahasiswa{}
	buku := model.Buku{}

	dataPeminjamanID := c.Param("id")

	if err := database.DB.First(&dataPeminjaman, dataPeminjamanID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Model(&dataPeminjaman).Update("tanggal_pengembalian", time.Now()).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.First(&mahasiswa, dataPeminjaman.MahasiswaID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mahasiswa.TotalPinjaman--

	database.DB.Save(&mahasiswa)

	if err := database.DB.First(&buku, dataPeminjaman.BukuID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	buku.Kuantitas++

	database.DB.Save(&buku)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success mengembalikan buku",
		"status":  201,
	})
}
