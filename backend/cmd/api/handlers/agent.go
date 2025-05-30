package handlers

import (
	"Arbitrax/pkg/output"
	agent_repo "Arbitrax/pkg/repositories/agent"
	exchanges_repo "Arbitrax/pkg/repositories/exchanges"
	strategy_repo "Arbitrax/pkg/repositories/strategies"

	"Arbitrax/pkg/services/validate"

	"encoding/json"
	"fmt"
	"net/http"
)

type AgentHandler struct {
	AgentRepo      agent_repo.Repository
	ExchangesRepo  exchanges_repo.Repository
	StrategiesRepo strategy_repo.Repository
}

func NewAgentHandler(repo agent_repo.Repository, e exchanges_repo.Repository, s strategy_repo.Repository) *AgentHandler {
	return &AgentHandler{
		AgentRepo:      repo,
		ExchangesRepo:  e,
		StrategiesRepo: s,
	}
}

// Response types
type CreateAgentResp struct {
	Agent *agent_repo.Model `json:"agent"`
}

type GetAgentsResp struct {
	Agents []*agent_repo.Model `json:"agents"`
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

	if r.CapitalAllocation <= 0 || r.CapitalAllocation > 1 {
		return fmt.Errorf("capital_allocation must be between 0 and 1")
	}

	if r.StopLoss < 0 || r.StopLoss > 1 {
		return fmt.Errorf("stop_loss must be between 0 and 1")
	}

	if r.TakeProfit < 0 || r.TakeProfit > 1 {
		return fmt.Errorf("take_profit must be between 0 and 1")
	}

	return nil
}

func (h *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	user, err := GetUserFromCtx(r)

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

	// If the agent is enabled, validate capital allocation sum
	if body.Enabled {
		// Get all existing agents for this user
		agents, err := h.AgentRepo.GetAllByUserUUID(r.Context(), user.UUID)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("failed to get user agents: %w", err)
		}

		// Calculate total capital allocation for enabled agents
		totalAllocation := body.CapitalAllocation
		for _, agent := range agents {
			if agent.Enabled {
				totalAllocation += agent.CapitalAllocation
			}
		}

		// Check if total allocation exceeds 1
		if totalAllocation > 1 {
			return http.StatusBadRequest, fmt.Errorf("total capital allocation for enabled agents would exceed 1.0 (current: %.2f, requested: %.2f, total: %.2f)",
				totalAllocation-body.CapitalAllocation, body.CapitalAllocation, totalAllocation)
		}
	}

	// Fetch all strategies and exchanges for validation
	allStrategies, err := h.StrategiesRepo.GetAll(r.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to fetch strategies: %w", err)
	}

	allExchanges, err := h.ExchangesRepo.GetAll(r.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to fetch exchanges: %w", err)
	}

	// Create maps for O(1) lookup
	strategyMap := make(map[string]bool)
	for _, strategy := range allStrategies {
		strategyMap[strategy.Key] = true
	}

	exchangeMap := make(map[string]bool)
	for _, exchange := range allExchanges {
		exchangeMap[exchange.Key] = true
	}

	// Validate strategies exist
	for _, strategyKey := range body.Strategies {
		if !strategyMap[strategyKey] {
			return http.StatusBadRequest, fmt.Errorf("strategy '%s' does not exist", strategyKey)
		}
	}

	// Validate exchanges exist
	for _, exchangeKey := range body.Exchanges {
		if !exchangeMap[exchangeKey] {
			return http.StatusBadRequest, fmt.Errorf("exchange '%s' does not exist", exchangeKey)
		}
	}

	// Create the agent
	agent, err := h.AgentRepo.Create(
		r.Context(),
		user.UUID,
		body.Name,
		body.Avatar,
		body.Enabled,
		body.CapitalAllocation,
		body.StopLoss,
		body.TakeProfit,
		body.Exchanges,
		body.Strategies,
		body.TestMode,
	)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to create agent: %w", err)
	}

	return output.SuccessResponse(w, r, &CreateAgentResp{
		Agent: agent,
	})
}

func (h *AgentHandler) GetAgents(w http.ResponseWriter, r *http.Request) (int, error) {
	user, err := GetUserFromCtx(r)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	agents, err := h.AgentRepo.GetAllByUserUUID(r.Context(), user.UUID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to get agents: %w", err)
	}

	return output.SuccessResponse(w, r, &GetAgentsResp{
		Agents: agents,
	})
}
