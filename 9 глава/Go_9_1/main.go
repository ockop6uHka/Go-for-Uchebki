package main

import (
	"fmt"
	"sync"
	"time"
)

type Withdrawal struct {
	amount  int
	success chan bool
}

type Account struct {
	balance     int
	deposits    chan int
	balances    chan int
	withdrawals chan Withdrawal
}

func NewAccount() *Account {
	account := &Account{
		deposits:    make(chan int),
		balances:    make(chan int),
		withdrawals: make(chan Withdrawal),
	}

	go account.teller()

	return account
}

func (a *Account) Deposit(amount int) {
	a.deposits <- amount
}

func (a *Account) Balance() int {
	return <-a.balances
}

func (a *Account) Withdraw(amount int) bool {
	ch := make(chan bool)
	a.withdrawals <- Withdrawal{amount, ch}
	return <-ch
}

func (a *Account) teller() {
	for {
		select {
		case amount := <-a.deposits:
			a.balance += amount
		case w := <-a.withdrawals:
			if w.amount > a.balance {
				w.success <- false
				continue
			}
			a.balance -= w.amount
			w.success <- true
		case a.balances <- a.balance:
		}
	}
}

func main() {
	account := NewAccount()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// Просто некоторая логика для демонстрации
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("Main goroutine is running...")
		}
		wg.Done()
	}()

	// Просто демонстрация использования аккаунта
	account.Deposit(100)
	fmt.Println("Balance after deposit:", account.Balance())

	withdrawSuccess := account.Withdraw(50)
	if withdrawSuccess {
		fmt.Println("Withdrawal successful. New balance:", account.Balance())
	} else {
		fmt.Println("Withdrawal failed. Insufficient funds.")
	}

	// Ожидание завершения работы главной горутины
	wg.Wait()
}
