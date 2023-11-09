package main

import (
	"errors"
	"fmt"
	"sync"
)

type InMemoryDatabase struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		data: make(map[string]string),
	}
}

func (db *InMemoryDatabase) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func (db *InMemoryDatabase) Get(key string) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, ok := db.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (db *InMemoryDatabase) Update(key, newValue string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, ok := db.data[key]
	if !ok {
		return errors.New("key not found")
	}
	db.data[key] = newValue
	return nil
}

func (db *InMemoryDatabase) Delete(key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, ok := db.data[key]
	if !ok {
		return errors.New("key not found")
	}
	delete(db.data, key)
	return nil
}

func main() {
	db := NewInMemoryDatabase()

	// Inserting key-value pairs
	db.Set("name", "Alice")
	db.Set("age", "30")

	// Getting values for keys
	name, err := db.Get("name")
	if err == nil {
		fmt.Println("Name:", name) // Output: Name: Alice
	}

	age, err := db.Get("age")
	if err == nil {
		fmt.Println("Age:", age) // Output: Age: 30
	}

	// Updating a key-value pair
	err = db.Update("age", "31")
	if err == nil {
		newAge, _ := db.Get("age")
		fmt.Println("Updated Age:", newAge) // Output: Updated Age: 31
	}

	// Deleting a key-value pair
	err = db.Delete("name")
	if err == nil {
		_, err := db.Get("name")
		if err != nil {
			fmt.Println("Name key deleted successfully") // Output: Name key deleted successfully
		}
	}
}
