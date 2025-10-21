package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Define the structure for Student
type Student struct {
	Name    string `json:"name"`
	Roll    string `json:"roll"`
	Address string `json:"address"`
}

// Define the student list and mutex for concurrency control
var (
	students []Student
	mu       sync.Mutex
)

// Get all students (GET /students)
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	if students == nil {
		students = make([]Student, 0) // Initialize if nil
	}
	json.NewEncoder(w).Encode(students)
}

// Add a new student (POST /add-student)
func AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mu.Lock()
	students = append(students, student)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// Delete a student (DELETE /delete-student?roll=123)
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roll := r.URL.Query().Get("roll")
	if roll == "" {
		http.Error(w, "Roll parameter missing", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, s := range students {
		if s.Roll == roll {
			students = append(students[:i], students[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted"})
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

// Update a student (PUT /update-student)
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updated Student
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, s := range students {
		if s.Roll == updated.Roll {
			students[i] = updated
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}
