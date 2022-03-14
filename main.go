package main

import (
	"fmt"
	"github.com/ayeniblessing101/loan/loan_matcher"
)

func main() {
	c := loan_matcher.Customer{
		Name:     "Erikaya",
		Cpf:      "123.456.789-10",
		Age:      29,
		Location: "SP",
		Income:   4000,
	}
	res := loan_matcher.GetLoanForCustomer(c)

	fmt.Printf("%+v", res)
}
