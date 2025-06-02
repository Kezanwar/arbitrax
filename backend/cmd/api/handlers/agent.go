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
	a agent_repo.Repository
	e exchanges_repo.Repository
	s strategy_repo.Repository
}

func NewAgentHandler(
	a agent_repo.Repository,
	e exchanges_repo.Repository,
	s strategy_repo.Repository,
) *AgentHandler {
	return &AgentHandler{
		a: a,
		e: e,
		s: s,
	}
}

// Response types
type CreateAgentResp struct {
	Agent *agent_repo.Model `json:"agent"`
}

type EditAgentResp struct {
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
	AiOrchestrated    bool     `json:"ai_orchestrated"`
	Enabled           bool     `json:"enabled"`
	TestMode          bool     `json:"test_mode"`
}

type EditAgentReqBody struct {
	Name              string   `json:"name" db:"name"`
	Avatar            string   `json:"avatar" db:"avatar"`
	CapitalAllocation float64  `json:"capital_allocation"`
	StopLoss          float64  `json:"stop_loss"`
	TakeProfit        float64  `json:"take_profit"`
	Exchanges         []string `json:"exchanges"`
	Strategies        []string `json:"strategies" db:"strategies"`
	AiOrchestrated    bool     `json:"ai_orchestrated"`
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

	if len(r.Exchanges) != 1 {
		return fmt.Errorf("one exchange must be specified")
	}

	for _, exc := range r.Exchanges {
		if !validate.StrNotEmpty(exc) {
			return fmt.Errorf("exchange name cannot be empty")
		}
	}

	if len(r.Strategies) == 0 {
		return fmt.Errorf("at least one strategy must be specified")
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

func (r *EditAgentReqBody) validate() error {

	if r.Name == "" {
		return fmt.Errorf("name must be specified")
	}

	if r.Avatar == "" {
		return fmt.Errorf("avatar must be specified")
	}

	if len(r.Exchanges) != 1 {
		return fmt.Errorf("one exchange must be specified")
	}

	for _, exc := range r.Exchanges {
		if !validate.StrNotEmpty(exc) {
			return fmt.Errorf("exchange name cannot be empty")
		}
	}

	if len(r.Strategies) == 0 {
		return fmt.Errorf("at least one strategy must be specified")
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

	// Get all existing agents for this user
	agents, err := h.a.GetAllByUserUUID(r.Context(), user.UUID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf(
			"failed to get user agents: %w",
			err,
		)
	}

	// Check if agent name already exists
	for _, agent := range agents {
		if agent.Name == body.Name {
			return http.StatusBadRequest, fmt.Errorf(
				"agent with name '%s' already exists",
				body.Name,
			)
		}
	}

	// If the agent is enabled, validate capital allocation sum
	if body.Enabled {
		// Calculate total capital allocation for enabled agents
		totalAllocation := body.CapitalAllocation
		for _, agent := range agents {
			if agent.Enabled {
				totalAllocation += agent.CapitalAllocation
			}
		}

		// Check if total allocation exceeds 1
		if totalAllocation > 1 {
			return http.StatusBadRequest, fmt.Errorf(
				"total capital allocation for enabled agents would exceed 1.0 (current: %.2f, requested: %.2f, total: %.2f)",
				totalAllocation-body.CapitalAllocation,
				body.CapitalAllocation,
				totalAllocation,
			)
		}
	}

	// Fetch all strategies for validation
	allStrategies, err := h.s.GetAll(r.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to fetch strategies: %w", err)
	}

	// Create map for O(1) lookup
	strategyMap := make(map[string]bool)
	for _, strategy := range allStrategies {
		strategyMap[strategy.Key] = true
	}

	// Validate strategies exist
	for _, strategyKey := range body.Strategies {
		if !strategyMap[strategyKey] {
			return http.StatusBadRequest, fmt.Errorf(
				"strategy '%s' does not exist",
				strategyKey,
			)
		}
	}

	// TODO: Uncomment when multiple exchanges are supported
	// allExchanges, err := h.ExchangesRepo.GetAll(r.Context())
	// if err != nil {
	// 	return http.StatusInternalServerError, fmt.Errorf("failed to fetch exchanges: %w", err)
	// }

	// TODO: Uncomment when multiple exchanges are supported
	// exchangeMap := make(map[string]bool)
	// for _, exchange := range allExchanges {
	// 	exchangeMap[exchange.Key] = true
	// }

	// TODO: Uncomment when multiple exchanges are supported
	// // Validate exchanges exist
	// for _, exchangeKey := range body.Exchanges {
	// 	if !exchangeMap[exchangeKey] {
	// 		return http.StatusBadRequest, fmt.Errorf("exchange '%s' does not exist", exchangeKey)
	// 	}
	// }

	// TODO: For now, hardcode exchanges to only allow 'ibkr' remove when supporting multiple exchanges
	body.Exchanges = []string{"ibkr"}

	// Create the agent
	agent, err := h.a.Create(
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
		body.AiOrchestrated,
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

	agents, err := h.a.GetAllByUserUUID(r.Context(), user.UUID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to get agents: %w", err)
	}

	return output.SuccessResponse(w, r, &GetAgentsResp{
		Agents: agents,
	})
}

func (h *AgentHandler) EditAgent(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	user, err := GetUserFromCtx(r)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	// Get agent UUID from URL params (assuming it's passed in the URL)
	agentUUID := r.URL.Query().Get("uuid")
	if agentUUID == "" {
		return http.StatusBadRequest, fmt.Errorf("agent uuid is required")
	}

	// Get the existing agent to verify ownership
	existingAgent, err := h.a.GetByUUID(r.Context(), agentUUID)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("agent not found")
	}

	// Verify the agent belongs to the user
	if existingAgent.UserUUID != user.UUID {
		return http.StatusForbidden, fmt.Errorf("you don't have permission to edit this agent")
	}

	var body EditAgentReqBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return http.StatusBadRequest, err
	}

	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	// Get all existing agents for this user
	agents, err := h.a.GetAllByUserUUID(r.Context(), user.UUID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf(
			"failed to get user agents: %w",
			err,
		)
	}

	// Check if agent name already exists (excluding current agent)
	for _, agent := range agents {
		if agent.Name == body.Name && agent.UUID != agentUUID {
			return http.StatusBadRequest, fmt.Errorf(
				"agent with name '%s' already exists",
				body.Name,
			)
		}
	}

	// If the agent is enabled, validate capital allocation sum
	if body.Enabled {
		// Calculate total capital allocation for enabled agents (excluding current agent)
		totalAllocation := body.CapitalAllocation
		for _, agent := range agents {
			if agent.Enabled && agent.UUID != agentUUID {
				totalAllocation += agent.CapitalAllocation
			}
		}

		// Check if total allocation exceeds 1
		if totalAllocation > 1 {
			return http.StatusBadRequest, fmt.Errorf(
				"total capital allocation for enabled agents would exceed 1.0 (current: %.2f, requested: %.2f, total: %.2f)",
				totalAllocation-body.CapitalAllocation,
				body.CapitalAllocation,
				totalAllocation,
			)
		}
	}

	// Fetch all strategies for validation
	allStrategies, err := h.s.GetAll(r.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to fetch strategies: %w", err)
	}

	// Create map for O(1) lookup
	strategyMap := make(map[string]bool)
	for _, strategy := range allStrategies {
		strategyMap[strategy.Key] = true
	}

	// Validate strategies exist
	for _, strategyKey := range body.Strategies {
		if !strategyMap[strategyKey] {
			return http.StatusBadRequest, fmt.Errorf(
				"strategy '%s' does not exist",
				strategyKey,
			)
		}
	}

	// TODO: Uncomment when multiple exchanges are supported
	// allExchanges, err := h.ExchangesRepo.GetAll(r.Context())
	// if err != nil {
	// 	return http.StatusInternalServerError, fmt.Errorf("failed to fetch exchanges: %w", err)
	// }

	// TODO: Uncomment when multiple exchanges are supported
	// exchangeMap := make(map[string]bool)
	// for _, exchange := range allExchanges {
	// 	exchangeMap[exchange.Key] = true
	// }

	// TODO: Uncomment when multiple exchanges are supported
	// // Validate exchanges exist
	// for _, exchangeKey := range body.Exchanges {
	// 	if !exchangeMap[exchangeKey] {
	// 		return http.StatusBadRequest, fmt.Errorf("exchange '%s' does not exist", exchangeKey)
	// 	}
	// }

	// TODO: For now, hardcode exchanges to only allow 'ibkr' remove when supporting multiple exchanges
	body.Exchanges = []string{"ibkr"}

	// Update the agent
	agent, err := h.a.Update(
		r.Context(),
		agentUUID,
		body.Name,
		body.Avatar,
		body.Enabled,
		body.CapitalAllocation,
		body.StopLoss,
		body.TakeProfit,
		body.Exchanges,
		body.Strategies,
		body.AiOrchestrated,
		body.TestMode,
	)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update agent: %w", err)
	}

	return output.SuccessResponse(w, r, &EditAgentResp{
		Agent: agent,
	})
}
