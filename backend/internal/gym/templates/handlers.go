package templates

import (
	"encoding/json"
	"errors"
	"life_grid/internal/gym"
	"life_grid/internal/gym/domain"
	"net/http"

	"github.com/google/uuid"
)

type TemplateHandler struct {
	service TemplateService
}

type TemplateExerciseHandler struct {
	service TemplateExerciseService
}

func NewTemplateHandler(service TemplateService) *TemplateHandler {
	return &TemplateHandler{service: service}
}

func NewTemplateExerciseHandler(service TemplateExerciseService) *TemplateExerciseHandler {
	return &TemplateExerciseHandler{service: service}
}

type CreateTemplateRequest struct {
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

type UpdateTemplateRequest struct {
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

type CreateTemplateExerciseRequest struct {
	ExerciseID   string  `json:"exercise_id"`
	OrderIndex   int     `json:"order_index"`
	TargetSets   int     `json:"target_sets"`
	TargetReps   int     `json:"target_reps"`
	TargetWeight float64 `json:"target_weight"`
}

type UpdateTemplateExerciseRequest struct {
	OrderIndex   int     `json:"order_index"`
	TargetSets   int     `json:"target_sets"`
	TargetReps   int     `json:"target_reps"`
	TargetWeight float64 `json:"target_weight"`
}

func (h *TemplateHandler) GetTemplateByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	template, err := h.service.GetTemplateByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrTemplateNotFound) {
			gym.SendError(w, http.StatusNotFound, "template not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, template)
}

func (h *TemplateHandler) ListTemplatesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok || userID == uuid.Nil {
		gym.SendError(w, http.StatusBadRequest, "invalid user_id")
		return
	}

	templates, err := h.service.ListTemplatesByUserID(r.Context(), userID)
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, templates)
}

func (h *TemplateHandler) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok || userID == uuid.Nil {
		gym.SendError(w, http.StatusBadRequest, "invalid user_id")
		return
	}

	var req CreateTemplateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	template, err := h.service.CreateTemplate(r.Context(), userID, req.Name, req.Notes)
	if err != nil {
		switch {
		case err == domain.ErrNameRequired, err == domain.ErrUserIDRequired:
			gym.SendError(w, http.StatusBadRequest, err.Error())
		default:
			gym.SendError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	gym.SendJSON(w, http.StatusCreated, template)
}

func (h *TemplateHandler) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateTemplateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	template, err := h.service.UpdateTemplate(r.Context(), id, req.Name, req.Notes)
	if err != nil {
		if errors.Is(err, ErrTemplateNotFound) {
			gym.SendError(w, http.StatusNotFound, "template not found")
			return
		}
		if err == domain.ErrNameRequired {
			gym.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, template)
}

func (h *TemplateHandler) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.DeleteTemplate(r.Context(), id); err != nil {
		if errors.Is(err, ErrTemplateNotFound) {
			gym.SendError(w, http.StatusNotFound, "template not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TemplateExerciseHandler) GetTemplateExerciseByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	templateExercise, err := h.service.GetTemplateExerciseByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrTemplateExerciseNotFound) {
			gym.SendError(w, http.StatusNotFound, "template exercise not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, templateExercise)
}

func (h *TemplateExerciseHandler) ListTemplateExercisesByTemplateID(w http.ResponseWriter, r *http.Request) {
	templateID, err := uuid.Parse(r.PathValue("templateID"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid template_id")
		return
	}

	items, err := h.service.ListTemplateExercisesByTemplateID(r.Context(), templateID)
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, items)
}

func (h *TemplateExerciseHandler) CreateTemplateExercise(w http.ResponseWriter, r *http.Request) {
	templateID, err := uuid.Parse(r.PathValue("templateID"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid template_id")
		return
	}

	var req CreateTemplateExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	exerciseID, err := uuid.Parse(req.ExerciseID)
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid exercise_id")
		return
	}

	templateExercise, err := h.service.CreateTemplateExercise(r.Context(), templateID, exerciseID, req.OrderIndex, req.TargetSets, req.TargetReps, req.TargetWeight)
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	gym.SendJSON(w, http.StatusCreated, templateExercise)
}

func (h *TemplateExerciseHandler) UpdateTemplateExercise(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateTemplateExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	templateExercise, err := h.service.UpdateTemplateExercise(r.Context(), id, req.OrderIndex, req.TargetSets, req.TargetReps, req.TargetWeight)
	if err != nil {
		if errors.Is(err, ErrTemplateExerciseNotFound) {
			gym.SendError(w, http.StatusNotFound, "template exercise not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	gym.SendJSON(w, http.StatusOK, templateExercise)
}

func (h *TemplateExerciseHandler) DeleteTemplateExercise(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.DeleteTemplateExercise(r.Context(), id); err != nil {
		if errors.Is(err, ErrTemplateExerciseNotFound) {
			gym.SendError(w, http.StatusNotFound, "template exercise not found")
			return
		}
		gym.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
