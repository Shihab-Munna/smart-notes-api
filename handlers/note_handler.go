package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"smart_note/database"
	"smart_note/models"
	"smart_note/utils"
	"strings"
)

// NoteHandler handles HTTP requests for /notes (GET all and POST)
func NoteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotes(w, r)
	case http.MethodPost:
		createNote(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NoteByIDHandler handles HTTP requests for /notes/{id} (GET, PUT, DELETE)
func NoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getNoteByID(w, r, id)
	case http.MethodPut:
		updateNoteByID(w, r, id)
	case http.MethodDelete:
		deleteNoteByID(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getNotes retrieves all notes from the database using GORM
func getNotes(w http.ResponseWriter, _ *http.Request) {
	var notes []models.Note
	if err := database.DB.Find(&notes).Error; err != nil {
		http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, notes)
}

// createNote adds a new note to the database using GORM
func createNote(w http.ResponseWriter, r *http.Request) {
	var newNote models.Note

	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&newNote).Error; err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, newNote)
}

// getNoteByID retrieves a specific note by ID from the database using GORM
func getNoteByID(w http.ResponseWriter, _ *http.Request, id uuid.UUID) {
	var note models.Note
	if err := database.DB.First(&note, "id = ?", id).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, note)
}

// updateNoteByID updates an existing note in the database using GORM
func updateNoteByID(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var existingNote models.Note
	if err := database.DB.First(&existingNote, "id = ?", id).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	existingNote.Title = updatedNote.Title
	existingNote.Content = updatedNote.Content
	if err := database.DB.Save(&existingNote).Error; err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, existingNote)
}

// deleteNoteByID deletes a note by ID from the database using GORM
func deleteNoteByID(w http.ResponseWriter, _ *http.Request, id uuid.UUID) {
	if err := database.DB.Delete(&models.Note{}, "id = ?", id).Error; err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
