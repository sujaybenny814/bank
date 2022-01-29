package validation

type AddBankModel struct {
	BankName   string `json:"bank_name" validate:"required,max=20,min=3"`
	IfscCode   string `json:"ifsc_code" validate:"required,max=20,min=10"`
	BranchName string `json:"branch_name" validate:"required,max=20,min=3"`
	BankId     string `json:"bank_id" validate:"required,max=20,min=6"`
}
