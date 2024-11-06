package models

import "sync"

// Student struct represents a student record.
type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

// In-memory map to hold student data.
var (
    StudentsData = make(map[int]Student) // The map where students are stored
    NextID       = 1                      // Counter for new student IDs
    DataLock     sync.Mutex               // Mutex for safe concurrent access
)
