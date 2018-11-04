package sample

import "fmt"

type Account struct {
	Name   string
	Income int
}

func ChangeLive() (account *Account) {
	WrongCase(account)
	fmt.Println("wrong:", account)
	account = &Account{}
	RightCase(account)
	fmt.Println("right:", account)
	return account
}

//don't use like this method
func WrongCase(account *Account) {
	account = &Account{
		Name:   "xiao",
		Income: 10,
	}
}

//don't use like this method
func RightCase(account *Account) {
	account.Name = "xiao"
	account.Income = 1000
}
