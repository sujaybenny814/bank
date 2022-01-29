package model

type Bank struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	BankName   string `json:"bank_name" validate:"required"`
	IfscCode   string `json:"ifsc_code" validate:"required"`
	BranchName string `json:"branch_name" validate:"required"`
	BankId     string `json:"bank_id" gorm:"unique" validate:"required"`
	Status     string `json:"status" gorm:"default:active"`
}
