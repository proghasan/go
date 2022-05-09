package fund

type Fund struct {
	balance int // balance is unexported ( Private ), because it's lowercase
}

// a regular function retuning a pointer to a Fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Fund{
		balance: initialBalance,
	}
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
