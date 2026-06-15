package repository

import (
	"sync"
	"groupie-tracker/internal/models"
)

type DataRepository struct {
	// RWMutex protects our cache from concurrent read/write race conditions
	mu        sync.RWMutex
	Artists   []models.Artist
	Locations map[int][]string
	Dates     map[int][]string
	Relations map[int]map[string][]string
}