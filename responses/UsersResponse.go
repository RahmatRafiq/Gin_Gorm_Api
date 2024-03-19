package responses

type UserResponses struct {
	ID          *uint   `json:"id"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	Role        *string `json:"role"`
	MahasiswaID *int    `json:"id_mahasiswa"`
}
