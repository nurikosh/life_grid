package session_exercises

import (
	"encoding/json"
	"errors"
	"life_grid/internal/gym"
	"life_grid/internal/gym/domain"
	"net/http"

	"github.com/google/uuid"
)

type SessionExerciseHandler struct {
	service SessionExerciseService
}

func NewSessionExerciseHandler(service SessionExerciseService) *SessionExerciseHandler {
	return &SessionExerciseHandler{service: service}
}

// DTOs
type CreateSessionExerciseRequest struct {
	ExerciseID string `json:"exercise_id"`
}

type SessionExerciseResponse struct {
	ID         string `json:"id"`
	SessionID  string `json:"session_id"`
	ExerciseID string `json:"exercise_id"`
}

// GET /sessions/{sessionID}/exercises/{id}
func (h *SessionExerciseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	se, err := h.service.GetSessionExerciseByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrSessionExerciseNotFound) {
			gym.SendError(w, http.StatusNotFound, "session exercise not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, toResponse(se))
}

// GET /sessions/{sessionID}/exercises
func (h *SessionExerciseHandler) ListBySessionID(w http.ResponseWriter, r *http.Request) {
	sessionID, err := uuid.Parse(r.PathValue("sessionID"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid session_id")
		return
	}

	items, err := h.service.ListSessionExercisesBySessionID(r.Context(), sessionID)
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := make([]SessionExerciseResponse, 0, len(items))
	for _, se := range items {
		resp = append(resp, toResponse(se))
	}

	gym.SendJSON(w, http.StatusOK, resp)
}

// POST /sessions/{sessionID}/exercises
func (h *SessionExerciseHandler) Create(w http.ResponseWriter, r *http.Request) {
	sessionID, err := uuid.Parse(r.PathValue("sessionID"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid session_id")
		return
	}

	var req CreateSessionExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	exerciseID, err := uuid.Parse(req.ExerciseID)
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid exercise_id")
		return
	}

	se, err := h.service.CreateSessionExercise(r.Context(), sessionID, exerciseID)
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	gym.SendJSON(w, http.StatusCreated, toResponse(se))
}

// DELETE /sessions/{sessionID}/exercises/{id}
func (h *SessionExerciseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.DeleteSessionExercise(r.Context(), id); err != nil {
		if errors.Is(err, ErrSessionExerciseNotFound) {
			gym.SendError(w, http.StatusNotFound, "session exercise not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func toResponse(se *domain.SessionExercise) SessionExerciseResponse {
	return SessionExerciseResponse{
		ID:         se.ID.String(),
		SessionID:  se.SessionID.String(),
		ExerciseID: se.ExerciseID.String(),
	}
}
