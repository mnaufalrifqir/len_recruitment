package model

import (
	"time"

	"gorm.io/gorm"
)

type DataPeminjaman struct {
	gorm.Model
	MahasiswaID         uint      `json:"mahasiswa_id"`
	BukuID              uint      `json:"buku_id"`
	TanggalPeminjaman   time.Time `json:"tanggal_peminjaman" gorm:"default:current_timestamp"`
	TanggalBatas        time.Time `json:"tanggal_batas"`
	TanggalPengembalian time.Time `json:"tanggal_pengembalian" gorm:"default:null"`
	Mahasiswa           Mahasiswa `json:"-" gorm:"foreignkey:MahasiswaID"`
	Buku                Buku      `json:"-" gorm:"foreignkey:BukuID"`
}
