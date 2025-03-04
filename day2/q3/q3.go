package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	balance int
	mu      sync.Mutex
}

func (a *account) deposit(money int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += money
	fmt.Printf("Deposited Rs.%d, New Balance: Rs.%d\n", money, a.balance)
}

func (a *account) withdraw(money int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < money {
		fmt.Println("Withdrawal failed: Insufficient balance")
		return
	}
	a.balance -= money
	fmt.Printf("Withdrawn Rs.%d, New Balance: Rs.%d\n", money, a.balance)
}

func main() {
	a := account{balance: 500}
	var wg sync.WaitGroup

	wg.Add(3)
	go a.deposit(400, &wg)
	go a.withdraw(1200, &wg)
	go a.deposit(100, &wg)

	wg.Wait()
	time.Sleep(time.Millisecond * 100)

	fmt.Println("Final Balance:", a.balance)
}
