// source - https://golangbyexample.com/mediator-design-pattern-golang/
// Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects.
// The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.

package behavioural

import (
	"fmt"
	"sync"
)

type train interface {
	requestArrival()
	departure()
	permitArrival()
}

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *passengerTrain) departure() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (g *goodsTrain) departure() {
	g.mediator.notifyFree()
	fmt.Println("GoodsTrain: Leaving")
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}

type mediator interface {
	canLand(train) bool
	notifyFree()
}

type stationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []train
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
