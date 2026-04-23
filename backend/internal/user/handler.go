package user

import (
	"encoding/json"
	"errors"
	"life_grid/internal/gym"
	"life_grid/internal/user/domain"
	"net/http"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

type registeterRequest struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	FullName string  `json:"full_name"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
}

type registerResponse struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	FullName string  `json:"full_name"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
}

type loginRequest struct {
	Email    string
	Password string
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registeterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		SendError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, err := h.service.Register(r.Context(), req.Email, req.Password, req.FullName, req.Weight, req.Height)

	if err != nil {
		switch {
		case err == domain.ErrEmailRequired, err == domain.ErrPasswordRequired:
			SendError(w, http.StatusBadRequest, err.Error())
		case err == domain.ErrEmailInvalid:
			SendError(w, http.StatusBadRequest, err.Error())
		case err == ErrEmailExists:
			SendError(w, http.StatusConflict, err.Error())
		default:
			SendError(w, http.StatusInternalServerError, "failed to create user")
		}
		return
	}

	resp := registerResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		FullName: user.FullName,
		Weight:   user.Weight,
		Height:   user.Height,
	}

	SendJSON(w, http.StatusCreated, resp)

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Email == "" || req.Password == "" {
		SendError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	token, err := h.service.Login(r.Context(), req.Email, req.Password)

	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			gym.SendError(w, http.StatusUnauthorized, "invalid credentials")
			return
		}

		gym.SendError(w, http.StatusInternalServerError, "failed to login")
		return
	}

	SendJSON(w, http.StatusOK, map[string]string{
		"access_token": token,
	})

}
