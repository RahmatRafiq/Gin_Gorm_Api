package models

// Import the package that contains the definition of the "Mahasiswa" struct

type Users struct {
	ID          *uint   `json:"id"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	Role        *string `json:"role"`
	MahasiswaID *int    `json:"id_mahasiswa" gorm:"column:id_mahasiswa"`
}
