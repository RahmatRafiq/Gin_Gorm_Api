package models

type Users struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	MahasiswaID uint   `json:"id_mahasiswa"`

	// Mahasiswa Mahasiswa `gorm:"foreignKey:MahasiswaID"`
}
