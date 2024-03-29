package main

import "fmt"

type employee interface{
	getName() string
	getSalary() int
}

type contractor struct{
	name string
	hourlyPay int
	hoursPerYear int
}

func(c contractor) getName() string{
	return c.name
}

func (c contractor) getSalary() int{
	return c.hourlyPay*c.hoursPerYear
}

type fullTime struct{
	name string
	salary int
}

func (ft fullTime) getSalary() int{
	return ft.salary
}

func (ft fullTime) getName() string{
	return ft.name
}

func test(e employee){
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("==============================")
}

func main(){
	test(fullTime{
		name:"Emanuel",
		salary: 400000,
	})
	test(contractor{
		name: "Carlos",
		hourlyPay: 80,
		hoursPerYear: 2080,
	})
}