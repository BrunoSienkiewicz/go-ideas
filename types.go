package main

import "math/rand"

type Idea struct {
	ID         int
	Name       string
	Category   string
	Attributes map[string]string
}

func NewIdea(name string, category string, attributes map[string]string) *Idea {
	return &Idea{
		ID:         rand.Intn(100),
		Name:       name,
		Category:   category,
		Attributes: attributes,
	}
}
