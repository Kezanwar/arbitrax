package handlers

import (
	options_memory_cache "Arbitrax/pkg/cache/options_memory"
	"Arbitrax/pkg/output"
	exchanges_repo "Arbitrax/pkg/repositories/exchanges"
	strategy_repo "Arbitrax/pkg/repositories/strategies"
	"net/http"
)

type OptionsHandler struct {
	s     strategy_repo.Repository
	e     exchanges_repo.Repository
	cache *options_memory_cache.Cache
}

func NewOptionsHandler(s strategy_repo.Repository, e exchanges_repo.Repository, cache *options_memory_cache.Cache) *OptionsHandler {
	return &OptionsHandler{
		s:     s,
		e:     e,
		cache: cache,
	}
}

type GetExchangesResp struct {
	Exchanges []*exchanges_repo.Model `json:"exchanges"`
}

func (h *OptionsHandler) GetExchangesOptions(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	opt, err := h.cache.GetExchanges(h.e.GetAll, r.Context())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return output.SuccessResponse(w, r, &GetExchangesResp{
		Exchanges: opt,
	})
}

type GetStrategiesResp struct {
	Strategies []*strategy_repo.Model `json:"strategies"`
}

func (h *OptionsHandler) GetStrategiesOptions(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	opt, err := h.cache.GetStrategies(h.s.GetAll, r.Context())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return output.SuccessResponse(w, r, &GetStrategiesResp{
		Strategies: opt,
	})
}
