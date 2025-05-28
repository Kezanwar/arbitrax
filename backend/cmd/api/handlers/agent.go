package handlers

import (
	"Arbitrax/pkg/output"
	agent_repo "Arbitrax/pkg/repositories/agent"

	"Arbitrax/pkg/services/validate"

	"encoding/json"
	"fmt"
	"net/http"
)

// Response types
type CreateAgentResp struct {
}

// Request types
type CreateAgentReqBody struct {
	Name              string   `json:"name" db:"name"`
	Avatar            string   `json:"avatar" db:"avatar"`
	CapitalAllocation float64  `json:"capital_allocation"`
	StopLoss          float64  `json:"stop_loss"`
	TakeProfit        float64  `json:"take_profit"`
	Exchanges         []string `json:"exchanges"`
	Strategies        []string `json:"strategies" db:"strategies"`
	Enabled           bool     `json:"enabled"`
	TestMode          bool     `json:"test_mode"`
}

func (r *CreateAgentReqBody) validate() error {

	if r.Name == "" {
		return fmt.Errorf("name must be specified")
	}

	if r.Avatar == "" {
		return fmt.Errorf("avatar must be specified")
	}

	if len(r.Exchanges) == 0 {
		return fmt.Errorf("at least one exchange must be specified")
	}

	for _, exc := range r.Exchanges {
		if !validate.StrNotEmpty(exc) {
			return fmt.Errorf("exchange name cannot be empty")
		}
	}

	if len(r.Strategies) == 0 {
		return fmt.Errorf("at least one exchange must be specified")
	}

	for _, exc := range r.Strategies {
		if !validate.StrNotEmpty(exc) {
			return fmt.Errorf("exchange name cannot be empty")
		}
	}

	if r.CapitalAllocation <= 0 {
		return fmt.Errorf("capital_allocation must be greater than 0")
	}

	if r.StopLoss < 0 || r.StopLoss > 1 {
		return fmt.Errorf("stop_loss must be between 0 and 1")
	}

	if r.TakeProfit < 0 || r.TakeProfit > 1 {
		return fmt.Errorf("take_profit must be between 0 and 1")
	}

	return nil
}

type AgentHandler struct {
	AgentRepo agent_repo.Repository
}

func NewAgentHandler(repo agent_repo.Repository) *AgentHandler {
	return &AgentHandler{AgentRepo: repo}
}

func (h *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	_, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	var body CreateAgentReqBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return http.StatusBadRequest, err
	}

	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	// exists, err := h.UserRepo.DoesEmailExist(r.Context(), body.Email)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }
	// if exists {
	// 	return http.StatusBadRequest, fmt.Errorf("This email already exists")
	// }

	// usr, err := h.UserRepo.Create(r.Context(), body.FirstName, body.LastName, body.Email, body.Password)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// tkn, err := jwt.Create(jwt.Keys.UUID, usr.UUID)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	return output.SuccessResponse(w, r, &output.MessageResponse{
		Message: "hello",
	})
}
