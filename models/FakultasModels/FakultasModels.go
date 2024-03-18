package models

type Fakultas struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	NamaFakultas string `json:"nama_fakultas"`
}
