package main

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Account
	Credits float64
}

func (a Account) String() string {
	return fmt.Sprintf("name:%v %v", a.LastName, a.FirstName)
}

func (a *Account) ChangeName(newAccount Account) {
	a.FirstName = newAccount.FirstName
	a.LastName = newAccount.LastName
}

func (e *Employee) AddCredits(num float64) {
	e.Credits += num
}

func (e *Employee) RemoveCredits(num float64) {
	e.Credits -= num
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func main() {
	shop := []Employee{
		Employee{Account: Account{FirstName: "zll", LastName: "zllsb"}, Credits: 100.0},
		Employee{Account: Account{FirstName: "baba", LastName: "baba"}, Credits: 100.0},
	}
	shop[0].AddCredits(10.0)
	shop[1].Account.ChangeName(Account{FirstName: "zllbaba", LastName: "zllbaba"})

	for _, item := range shop {
		fmt.Printf("%v's Credits is %v \n", item.Account, item.Credits)
	}
}
