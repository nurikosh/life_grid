package transport

import (
	"encoding/json"
	"gym/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ExerciseHandler struct {
	s service.ExerciseService
}

func NewExerciseHandler(s service.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{s: s}
}

func (e ExerciseHandler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	var req service.ExerciseInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	input := service.ExerciseInput{
		Name:        req.Name,
		MuscleGroup: req.MuscleGroup,
		Description: req.Description,
	}

	ex, err := e.s.Create(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	sendJSON(w, http.StatusOK, ex)
}

func (e ExerciseHandler) ListExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := e.s.List(r.Context())
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, http.StatusOK, exercises)
}

func (e ExerciseHandler) GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		sendError(w, "invalid id format", http.StatusBadRequest)
		return
	}
	ex, err := e.s.GetByID(r.Context(), id)
	if err != nil {
		sendError(w, err.Error(), http.StatusNotFound)
		return
	}

	sendJSON(w, http.StatusOK, ex)
}

func (e ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		sendError(w, "invalid id format", http.StatusBadRequest)
		return
	}

	if err := e.s.Delete(r.Context(), id); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
