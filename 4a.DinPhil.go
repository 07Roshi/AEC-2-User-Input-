package main

import (
"fmt"
"sync"
"time"
)

type Philosopher struct {
id int
leftFork, rightFork *sync.Mutex
diningCycles  int
}

type DiningTable struct {
philosophers []*Philosopher
waiter *sync.Mutex
}

func (p *Philosopher) think() {
fmt.Printf("Philosopher %d is thinking\n", p.id)
time.Sleep(time.Second)
}

func (p *Philosopher) eat() {
p.leftFork.Lock()
p.rightFork.Lock()

fmt.Printf("Philosopher %d is eating\n", p.id)
time.Sleep(time.Second)

p.rightFork.Unlock()
p.leftFork.Unlock()

p.diningCycles++
}

func (p *Philosopher) dine(table *DiningTable, maxDiningCycles int) {
for p.diningCycles < maxDiningCycles {
p.think()

table.waiter.Lock()

p.eat()
table.waiter.Unlock()
}
}

func main() {
var numPhilosophers, maxDiningCycles int

fmt.Print("Enter the number of philosophers: ")
fmt.Scan(&numPhilosophers)

fmt.Print("Enter the maximum dining cycles: ")
fmt.Scan(&maxDiningCycles)

table := &DiningTable{
philosophers: make([]*Philosopher, numPhilosophers),
waiter: &sync.Mutex{},
}

// Create forks
forks := make([]*sync.Mutex, numPhilosophers)
for i := 0; i < numPhilosophers; i++ {
forks[i] = &sync.Mutex{}
}

// Create philosophers and assign forks
for i := 0; i < numPhilosophers; i++ {
table.philosophers[i] = &Philosopher{
id: i,
leftFork:  forks[i],
rightFork: forks[(i+1)%numPhilosophers],
}
}

// Start dining
var wg sync.WaitGroup
wg.Add(numPhilosophers)
for i := 0; i < numPhilosophers; i++ {
go func(p *Philosopher) {
defer wg.Done()
p.dine(table, maxDiningCycles)
}(table.philosophers[i])
}
wg.Wait()
}


// Output
// Enter the number of philosophers: 3
// Enter the maximum dining cycles: 2
// Philosopher 2 is thinking
// Philosopher 0 is thinking
// Philosopher 1 is thinking
// Philosopher 1 is eating
// Philosopher 1 is thinking
// Philosopher 2 is eating
// Philosopher 2 is thinking
// Philosopher 1 is eating
// Philosopher 0 is eating
// Philosopher 0 is thinking
// Philosopher 2 is eating
// Philosopher 0 is eating