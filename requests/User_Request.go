package requests

type UserRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	MahasiswaID int    `json:"id_mahasiswa" binding:"required"`
	// Email    string `json:"email" binding:"required,email,unique"`
}
