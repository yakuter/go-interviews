package main

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if subscribers, ok := ps.subscribers[topic]; ok {
		for _, ch := range subscribers {
			go func(c chan string) {
				c <- message
			}(ch)
		}
	}
}

func main() {
	pubsub := NewPubSub()

	// Subscriber 1
	subscriber1 := pubsub.Subscribe("topic1")
	go func() {
		for {
			message := <-subscriber1
			fmt.Println("Subscriber 1 received message:", message)
		}
	}()

	// Subscriber 2
	subscriber2 := pubsub.Subscribe("topic2")
	go func() {
		for {
			message := <-subscriber2
			fmt.Println("Subscriber 2 received message:", message)
		}
	}()

	// Publisher
	pubsub.Publish("topic1", "Hello, topic1 subscribers!")
	pubsub.Publish("topic2", "Greetings, topic2 subscribers!")

	// Wait for a while to allow subscribers to process messages
	time.Sleep(3 * time.Second)
}
