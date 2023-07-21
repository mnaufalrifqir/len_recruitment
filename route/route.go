package route

import (
	"len_recruitment/controller"

	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo{
	e := echo.New()

	//CRUD Mahasiswa
	e.POST("/mahasiswa", controller.CreateMahasiswa)
	e.GET("/mahasiswa", controller.GetAllMahasiswa)
	e.GET("/mahasiswa/:id", controller.GetMahasiswaByID)
	e.PUT("/mahasiswa/:id", controller.UpdateMahasiswaByID)
	e.DELETE("/mahasiswa/:id", controller.DeleteMahasiswaByID)

	//CRUD Buku
	e.POST("/buku", controller.CreateBuku)
	e.GET("/buku", controller.GetAllBuku)
	e.GET("/buku/:id", controller.GetBukuByID)
	e.PUT("/buku/:id", controller.UpdateBukuByID)
	e.DELETE("/buku/:id", controller.DeleteBukuByID)

	//CRUD Data Peminjaman
	e.POST("/peminjaman", controller.CreateDataPeminjaman)
	e.PUT("/peminjaman/:id", controller.PengembalianBuku)
	return e
}