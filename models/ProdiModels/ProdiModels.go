package models

type Prodi struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	NamaProdi  string `json:"nama_prodi"`
	FakultasID uint   `json:"id_fakultas"`
}
