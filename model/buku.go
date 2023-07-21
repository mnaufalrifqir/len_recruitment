package model

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Judul     string `json:"judul"`
	Penulis   string `json:"penulis"`
	Kuantitas uint    `json:"kuantitas"`
	Tempat    string `json:"tempat"`
	DataPeminjaman []DataPeminjaman `json:"-" gorm:"foreignkey:BukuID"`
}
