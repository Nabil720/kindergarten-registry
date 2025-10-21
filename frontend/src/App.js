import React, { useState, useEffect } from "react";
import StudentForm from "./components/StudentForm";
import StudentList from "./components/StudentList";

function App() {
  const [students, setStudents] = useState([]);

  const fetchStudents = () => {
    fetch("http://localhost:5000/students")
      .then((res) => {
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        return res.json();
      })
      .then((data) => setStudents(data || []))
      .catch((err) => console.error("Fetch error:", err));
  };

  useEffect(() => {
    fetchStudents();
  }, []);

  const addStudent = async (student) => {
    try {
      await fetch("http://localhost:5000/add-student", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(student),
      });
      fetchStudents(); // Refresh student list after adding
    } catch (err) {
      console.error("Error adding student:", err);
    }
  };

  const handleDelete = async (roll) => {
    if (window.confirm("Are you sure you want to delete this student?")) {
      try {
        await fetch(`http://localhost:5000/delete-student?roll=${roll}`, {
          method: "DELETE",
        });
        fetchStudents(); // Refresh list after deletion
      } catch (err) {
        console.error("Error deleting student:", err);
      }
    }
  };

  const handleEdit = async (student) => {
    try {
      await fetch(`http://localhost:5000/update-student`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(student),
      });
      fetchStudents(); // Refresh list after update
    } catch (err) {
      console.error("Error updating student:", err);
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <h1>Kindergarten School Registry</h1>
      <StudentForm onAddStudent={addStudent} />
      <StudentList
        students={students}
        onDelete={handleDelete}
        onEdit={handleEdit}
      />
    </div>
  );
}

export default App;
