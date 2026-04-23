package sets

import (
	"encoding/json"
	"errors"
	"life_grid/internal/gym"
	"net/http"

	"github.com/google/uuid"
)

type SetHandler struct {
	service SetService
}

func NewSetHandler(service SetService) *SetHandler {
	return &SetHandler{service: service}
}

type CreateSetRequest struct {
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
	Order  int     `json:"order_num"`
}

type UpdateSetRequest struct {
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
	Order  int     `json:"order_num"`
}

func (h *SetHandler) GetSetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	set, err := h.service.GetSetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrSetNotFound) {
			gym.SendError(w, http.StatusNotFound, "set not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, set)
}

func (h *SetHandler) ListSetsBySessionExerciseID(w http.ResponseWriter, r *http.Request) {
	sessionExerciseID, err := uuid.Parse(r.PathValue("sessionExerciseID"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid session_exercise_id")
		return
	}

	items, err := h.service.ListSetsBySessionExerciseID(r.Context(), sessionExerciseID)
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, items)
}

func (h *SetHandler) CreateSet(w http.ResponseWriter, r *http.Request) {
	sessionExerciseID, err := uuid.Parse(r.PathValue("sessionExerciseID"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid session_exercise_id")
		return
	}

	var req CreateSetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	set, err := h.service.CreateSet(r.Context(), sessionExerciseID, req.Reps, req.Weight, req.Order)
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	gym.SendJSON(w, http.StatusCreated, set)
}

func (h *SetHandler) UpdateSet(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateSetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	set, err := h.service.UpdateSet(r.Context(), id, req.Reps, req.Weight, req.Order)
	if err != nil {
		if errors.Is(err, ErrSetNotFound) {
			gym.SendError(w, http.StatusNotFound, "set not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, set)
}

func (h *SetHandler) DeleteSet(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.DeleteSet(r.Context(), id); err != nil {
		if errors.Is(err, ErrSetNotFound) {
			gym.SendError(w, http.StatusNotFound, "set not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
