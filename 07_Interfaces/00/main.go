package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

type Contract struct {
	empId    int
	basicPay int
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicPay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpenseToConpany(employees)
}

func totalExpenseToConpany(s []SalaryCalculator) {
	totalExpense := 0
	for _, v := range s {
		totalExpense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense/Month: %d", totalExpense)
}
