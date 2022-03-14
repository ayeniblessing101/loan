package loan_matcher

func GetLoanForCustomer(c Customer) Response {
	var loans []Loan
	if c.Income <= 3000 || (c.Income > 3000 && c.Income < 5000) {
		loan := Loan{LoanType: LoanTypePersonal}
		loans = append(loans, loan)
	}

	if c.Income >= 5000 {
		loan1 := Loan{LoanType: LoanTypePersonal}
		loan2 := Loan{LoanType: LoanTypePayroll}
		loans = append(loans, loan1, loan2)
	}

	if (c.Income <= 3000 && c.Age < 30 && c.Location == "SP") ||
		(c.Income > 3000 && c.Income < 5000) && c.Location == "SP" ||
		(c.Income >= 5000 && c.Age < 30) {
		loan := Loan{LoanType: LoanTypeCollateralized}
		loans = append(loans, loan)
	}
	return Response{
		Customer: c.Name,
		Loans:    loans,
	}

}
