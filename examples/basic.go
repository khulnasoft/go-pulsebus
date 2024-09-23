package main

import (
	"fmt"

	"github.com/khulnasoft/go-pulsebus"
)

type A struct {
	a string
	b int
}

type B struct {
	c int64
	d float64
}

func (b *B) Number() float64 {
	return b.d
}

type Number interface {
	Number() float64
}

func generate(bus *pulsebus.Bus) {
	bus.Publish(pulsebus.Event{
		Type:  "a",
		Value: A{"yeah", 4},
	})
	bus.Publish(pulsebus.Event{
		Type:  "b",
		Value: B{42, 5.5},
	})
	bus.Publish(pulsebus.Event{
		Type:  "n",
		Value: Number(&B{42, 5.5}),
	})
	bus.Close()
}

func main() {

	bus := pulsebus.NewBus()
	subscription := bus.Subscribe()

	go generate(bus)

	for event := range subscription.Events() {
		fmt.Printf("Event: %+v\n", event)
		switch event.Value.(type) {
		case string:
			fmt.Println("STRING", event.Value.(string))
		case Number:
			fmt.Println("Number", event.Value.(Number).Number())
		case A:
			fmt.Println("A type", event.Value.(A).a)
		case B:
			fmt.Println("B type", event.Value.(B).c)
		default:
			fmt.Println("blerg...")
		}
	}

}
