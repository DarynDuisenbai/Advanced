package main

type Person struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Address      Address   `json:"address"`
	Contacts     []Contact `json:"contacts"`
	IsStudent    bool      `json:"isStudent"`
	Grades       []int     `json:"grades"`
	RegisteredAt string    `json:"registeredAt"`
}
