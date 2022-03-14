package loan_matcher

type (
	LoanType string

	Customer struct {
		Name     string  `json:"name"`
		Cpf      string  `json:"cpf"`
		Age      int     `json:"age"`
		Location string  `json:"location"`
		Income   float64 `json:"income"`
	}

	Loan struct {
		LoanType LoanType `json:"type"`
	}

	Response struct {
		Customer string `json:"customer"`
		Loans    []Loan `json:"loans"`
	}
)

const (
	LoanTypePersonal       LoanType = "Personal"
	LoanTypeCollateralized LoanType = "Collateralized"
	LoanTypePayroll        LoanType = "Payroll"
)
