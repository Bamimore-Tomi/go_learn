package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depesiting %d to account with balance: %d\n", value, balance)

	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from acount with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("A mutex sample")

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
