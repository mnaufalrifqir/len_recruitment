package model

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama           string           `json:"nama"`
	NIM            string           `json:"nim"`
	Jurusan        string           `json:"jurusan"`
	TotalPinjaman  int              `json:"-" gorm:"default:0"`
	DataPeminjaman []DataPeminjaman `json:"-" gorm:"foreignkey:MahasiswaID"`
}
