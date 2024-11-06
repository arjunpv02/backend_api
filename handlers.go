package handlers

import (
    "encoding/json"
    "fmt"
    "models"
    "net/http"
    "strconv"
)

// Get all students
func GetStudents(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    models.DataLock.Lock()
    defer models.DataLock.Unlock()

    // Return all students from the in-memory map
    json.NewEncoder(w).Encode(models.StudentsData)
}

// Create a new student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
    var newStudent models.Student
    if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    models.DataLock.Lock()
    defer models.DataLock.Unlock()

    // Assign a new ID and save the student to the map
    newStudent.ID = models.NextID
    models.StudentsData[models.NextID] = newStudent
    models.NextID++ // Increment for next student

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newStudent)
}

// Get a student by ID
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    studentID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    models.DataLock.Lock()
    defer models.DataLock.Unlock()

    student, exists := models.StudentsData[studentID]
    if !exists {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(student)
}

// Update student information
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    studentID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var updatedStudent models.Student
    if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    models.DataLock.Lock()
    defer models.DataLock.Unlock()

    existingStudent, exists := models.StudentsData[studentID]
    if !exists {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    // Update fields
    existingStudent.Name = updatedStudent.Name
    existingStudent.Age = updatedStudent.Age
    existingStudent.Email = updatedStudent.Email

    models.StudentsData[studentID] = existingStudent
    json.NewEncoder(w).Encode(existingStudent)
}

// Delete student by ID
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    studentID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    models.DataLock.Lock()
    defer models.DataLock.Unlock()

    if _, exists := models.StudentsData[studentID]; exists {
        delete(models.StudentsData, studentID)
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "Student not found", http.StatusNotFound)
    }
}
