package models

type Mahasiswa struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Nama    string `json:"nama"`
	Nim     string `json:"nim"`
	ProdiID uint   `json:"id_prodi"`
}
