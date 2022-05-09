package main

func main() {
	CreateNewAccount("Hasan Sheikh", 123456, 200.00)
	// fmt.Println("Your current Balance:", Account.balance)
}

type Account struct {
	accountHolder string
	accountNumber int
	balance       float64
}

// create new account
func CreateNewAccount(name string, accountNumber int, initBalance float64) *Account {
	return &Account{
		accountHolder: name,
		accountNumber: accountNumber,
		balance:       initBalance,
	}
}

// current balance
func (a *Account) Balance() float64 {
	return a.balance
}
