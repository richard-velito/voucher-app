package models

type Voucher struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Label       string `json:"label"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Important   bool   `json:"important"`
}
