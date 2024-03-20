package models

type Fakultas struct {
	ID           *uint   `json:"id" gorm:"column:primary_key"`
	NamaFakultas *string `json:"column:nama_fakultas"`
}
