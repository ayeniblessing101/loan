package loan_matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLoanForCustomer(t *testing.T) {
	cases := []struct {
		in       Customer
		got      Response
		expected Response
	}{
		{
			in:       Customer{Name: "Blessing", Cpf: "", Location: "SP", Age: 29, Income: 3000},
			expected: Response{Customer: "Blessing", Loans: []Loan{{LoanType: "Personal"}, {LoanType: "Collateralized"}}},
		},
		{
			in:       Customer{Name: "Biden", Cpf: "", Location: "NL", Age: 30, Income: 3000},
			expected: Response{Customer: "Biden", Loans: []Loan{{LoanType: "Personal"}}},
		},
		{
			in:       Customer{Name: "Joe", Cpf: "", Location: "SP", Age: 35, Income: 4000},
			expected: Response{Customer: "Joe", Loans: []Loan{{LoanType: "Personal"}, {LoanType: "Collateralized"}}},
		},
		{
			in:       Customer{Name: "hopeful", Cpf: "", Location: "SP", Age: 29, Income: 5000},
			expected: Response{Customer: "hopeful", Loans: []Loan{{LoanType: "Personal"}, {LoanType: "Payroll"}, {LoanType: "Collateralized"}}},
		},
	}

	for _, c := range cases {
		c.got = GetLoanForCustomer(c.in)
		assert.Equal(t, c.got.Customer, c.expected.Customer)
		assert.Equal(t, c.got.Loans, c.expected.Loans)
		assert.NotNil(t, c.got)
		assert.Equal(t, c.got.Loans[0].LoanType, c.expected.Loans[0].LoanType)
	}
}
