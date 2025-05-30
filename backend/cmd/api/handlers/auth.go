package handlers

import (
	user_repo "Arbitrax/pkg/repositories/user"

	"Arbitrax/pkg/output"
	"Arbitrax/pkg/services/jwt"
	"Arbitrax/pkg/services/validate"

	"encoding/json"
	"fmt"
	"net/http"
)

// Response types
type ManualAuthResp struct {
	User  *user_repo.Model `json:"user"`
	Token string           `json:"token"`
}

type AutoAuthResp struct {
	User *user_repo.Model `json:"user"`
}

// Request types
type RegisterReqBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r *RegisterReqBody) validate() error {
	if !validate.StrNotEmpty(r.FirstName, r.LastName, r.Email, r.Password) {
		return fmt.Errorf("Request body invalid")
	}
	return nil
}

type SignInReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *SignInReqBody) validate() error {
	if !validate.StrNotEmpty(r.Email, r.Password) {
		return fmt.Errorf("Request body invalid")
	}
	return nil
}

type AuthHandler struct {
	UserRepo user_repo.Repository
}

func NewAuthHandler(repo user_repo.Repository) *AuthHandler {
	return &AuthHandler{UserRepo: repo}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	var body RegisterReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	exists, err := h.UserRepo.DoesEmailExist(r.Context(), body.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if exists {
		return http.StatusBadRequest, fmt.Errorf("This email already exists")
	}

	usr, err := h.UserRepo.Create(r.Context(), body.FirstName, body.LastName, body.Email, body.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	tkn, err := jwt.Create(jwt.Keys.UUID, usr.UUID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return output.SuccessResponse(w, r, &ManualAuthResp{
		User:  usr,
		Token: tkn,
	})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	var body SignInReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	usr, err := h.UserRepo.GetByEmail(r.Context(), body.Email)

	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("Invalid credentials")
	}

	if !usr.IsPassword(body.Password) {
		return http.StatusBadRequest, fmt.Errorf("Invalid credentials")
	}

	tkn, err := jwt.Create(jwt.Keys.UUID, usr.UUID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to create authorization session")
	}

	return output.SuccessResponse(w, r, &ManualAuthResp{
		User:  usr,
		Token: tkn,
	})
}

func (h *AuthHandler) Initialize(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	return output.SuccessResponse(w, r, &AutoAuthResp{
		User: usr,
	})
}
