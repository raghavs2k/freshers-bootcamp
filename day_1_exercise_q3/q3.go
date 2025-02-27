package main

import (
	"fmt"
)

type Employee interface {
	CalculateSalary() int
}

type FullTime struct {
	dailyWage  int
	daysWorked int
}

type Contractor struct {
	dailyWage  int
	daysWorked int
}

type Freelancer struct {
	hourlyWage  int
	hoursWorked int
}

func (f FullTime) CalculateSalary() int {
	return f.dailyWage * f.daysWorked
}

func (c Contractor) CalculateSalary() int {
	return c.dailyWage * c.daysWorked
}

func (fl Freelancer) CalculateSalary() int {
	return fl.hourlyWage * fl.hoursWorked
}

func main() {
	fullTimeEmployee := FullTime{dailyWage: 500, daysWorked: 30}
	contractor := Contractor{dailyWage: 150, daysWorked: 20}
	freelancer := Freelancer{hourlyWage: 100, hoursWorked: 160}

	employees := []Employee{fullTimeEmployee, contractor, freelancer}

	for _, employee := range employees {
		fmt.Printf("Salary: %d\n", employee.CalculateSalary())
	}
}
