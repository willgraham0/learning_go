package main

import (
	"fmt"
	"sync"
)

// ChopS is a chopstick object.
type ChopS struct{ sync.Mutex }

// Philo is a philosopher object.
type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

// eat makes the philosopher do the following 3 times: grab a left and
// right chopstick, eat and then release the chopsticks.
func (p Philo) eat(host chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		// Send flag that philosopher is eating to host.
		host <- true

		// Grab the chopsticks.
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", p.id)
		fmt.Println("finishing eating ", p.id)

		// Release the chopsticks.
		p.rightCS.Unlock()
		p.leftCS.Unlock()

		// Send flag that philosopher has finished eating to host.
		<-host
	}

	// Report as complete.
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var numberOfPhilosophers = 5
	var numberOfChopsticks = 5

	wg.Add(numberOfPhilosophers)

	// Create request communication channel of size 2.
	host := make(chan bool, 2)

	// Create chopsticks.
	CSticks := make([]*ChopS, numberOfChopsticks)
	for i := 0; i < numberOfChopsticks; i++ {
		CSticks[i] = new(ChopS)
	}

	// Create philosophers.
	philos := make([]*Philo, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	// Start the dinner.
	for i := 0; i < numberOfPhilosophers; i++ {
		go philos[i].eat(host, &wg)
	}

	wg.Wait()
}
