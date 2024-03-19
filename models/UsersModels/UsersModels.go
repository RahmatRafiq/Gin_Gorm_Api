package models

// Import the package that contains the definition of the "Mahasiswa" struct

type Users struct {
	ID          *uint   `json:"id" gorm:"column:id"`
	Username    *string `json:"username" gorm:"column:username"`
	Email       *string `json:"email" gorm:"column:email"`
	Password    *string `json:"password" gorm:"column:password"`
	Role        *string `json:"role" gorm:"column:role"`
	MahasiswaID *int    `json:"id_mahasiswa" gorm:"column:id_mahasiswa"`
}
