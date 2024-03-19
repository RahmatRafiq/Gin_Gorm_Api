package requests

type UserRequest struct {
	Username    string `json:"username" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Role        string `json:"role" form:"role" binding:"required"`
	MahasiswaID int    `json:"id_mahasiswa" form:"id_mahasiswa" binding:"required"`
	// Email    string `json:"email" binding:"required,email,unique"`
}
