package sessions

import (
	"life_grid/internal/gym"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type SessionHandler struct {
	service SessionService
}

func NewSessionHandler(service SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) GetSessionByID(w http.ResponseWriter, r *http.Request) {
	session, err := h.service.GetSessionByID(r.Context(), r.Context().Value("sessionID").(uuid.UUID))

	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to get session")
		return
	}

	gym.SendJSON(w, http.StatusOK, session)
}

func (h *SessionHandler) ListSessionByUserID(w http.ResponseWriter, r *http.Request) {
	sessions, err := h.service.ListSessionByUserID(r.Context(), r.Context().Value("userID").(uuid.UUID))
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to list sessions")
		return
	}
	gym.SendJSON(w, http.StatusOK, sessions)
}

func (h *SessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	session, err := h.service.StartSession(r.Context(), r.Context().Value("userID").(uuid.UUID), r.FormValue("name"))
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to start session")
		return
	}
	gym.SendJSON(w, http.StatusCreated, session)
}

func (h *SessionHandler) EndSession(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Context().Value("sessionID").(string))
	if err != nil {
		gym.SendError(w, http.StatusBadRequest, "invalid session id")
		return
	}
	session, err := h.service.EndSession(r.Context(), id, time.Now())
	if err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to end session")
		return
	}
	gym.SendJSON(w, http.StatusOK, session)
}

func (h *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteSession(r.Context(), r.Context().Value("sessionID").(uuid.UUID)); err != nil {
		gym.SendError(w, http.StatusInternalServerError, "Failed to delete session")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
