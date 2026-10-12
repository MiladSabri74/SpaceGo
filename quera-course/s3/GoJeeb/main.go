package main

type Account interface {
	Deposit(amount int) string
	Withdraw(amount int) string
	GetBalance() int
	CalculateProfit() int
}

// ------------ Structures ---------------------------------------------
type BaseAccount struct {
	balance int
}
type EverydayAccount struct {
	BaseAccount
}
type SavingsAccount struct {
	BaseAccount
}

type BusinessAccount struct {
	BaseAccount
}

func NewEverydayAccount() *EverydayAccount {
	ea := &EverydayAccount{}
	return ea
}

func NewSavingsAccount() *SavingsAccount {
	sa := &SavingsAccount{}
	return sa
}

func NewBusinessAccount() *BusinessAccount {
	ba := &BusinessAccount{}
	return ba
}
func (account *BaseAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Invalid amount"
	}
	account.balance += amount
	return "Success"
}

func (account *BaseAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Invalid amount"
	}
	if account.balance < amount {
		return "Insufficient funds"
	}
	account.balance -= amount
	return "Success"
}

func (account *BaseAccount) GetBalance() int {
	return account.balance
}

func (account *EverydayAccount) CalculateProfit() int {
	return 0
}

func (account *SavingsAccount) CalculateProfit() int {
	return int(account.balance * 5 / 100)
}

func (account *BusinessAccount) CalculateProfit() int {
	return int(account.balance * 2 / 100)
}

func Transfer(sender, receiver Account, amount int) string {
	if sender == nil || receiver == nil {
		return "Invalid account"
	}
	withdrawResult := sender.Withdraw(amount)
	if withdrawResult != "Success" {
		return withdrawResult
	}
	return receiver.Deposit(amount)
}
