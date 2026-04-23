package exercises

import (
	"encoding/json"
	"life_grid/internal/gym"
	"life_grid/internal/gym/domain"
	"net/http"

	"github.com/google/uuid"
)

type ExerciseHandler struct {
	exerciseService ExerciseService
}

func NewExerciseHandler(ExerciseService ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{exerciseService: ExerciseService}
}

type ExerciseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	MuscleGroup string `json:"muscle_group"`
}

func (h *ExerciseHandler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	var req ExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	exercise, err := h.exerciseService.CreateExercise(r.Context(), req.Name, req.Description, req.MuscleGroup)

	if err != nil {
		switch {
		case err == domain.ErrNameRequired, err == domain.ErrMuscleGroupRequired:
			gym.SendError(w, http.StatusBadRequest, err.Error())
		default:
			gym.SendError(w, http.StatusInternalServerError, "Failed to create exercise")
		}
		return
	}

	gym.SendJSON(w, http.StatusCreated, exercise)
}

func (h *ExerciseHandler) ListExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := h.exerciseService.ListExercises(r.Context())
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to list exercises")
		return
	}
	gym.SendJSON(w, http.StatusOK, exercises)
}

func (h *ExerciseHandler) GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	exercise, err := h.exerciseService.GetExerciseByID(r.Context(), r.Context().Value("exerciseID").(uuid.UUID))
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to get exercise")
		return
	}

	gym.SendJSON(w, http.StatusOK, exercise)
}

func (h *ExerciseHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	var req ExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	exercise, err := h.exerciseService.UpdateExercise(r.Context(), r.Context().Value("exerciseID").(uuid.UUID), req.Name, req.Description, req.MuscleGroup)

	if err != nil {
		switch {
		case err == domain.ErrNameRequired, err == domain.ErrMuscleGroupRequired:
			gym.SendError(w, http.StatusBadRequest, err.Error())
		default:
			gym.SendError(w, http.StatusInternalServerError, "Failed to update exercise")
		}
		return
	}

	gym.SendJSON(w, http.StatusOK, exercise)
}

func (h *ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	err := h.exerciseService.DeleteExercise(r.Context(), r.Context().Value("exerciseID").(uuid.UUID))
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to delete exercise")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
