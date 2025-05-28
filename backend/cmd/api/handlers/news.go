package handlers

import (
	user_repo "Arbitrax/pkg/repositories/user"
	"os"

	"Arbitrax/pkg/output"

	"fmt"
	"net/http"
)

// Response types
type NewsResp struct {
	User *user_repo.ToClient `json:"news"`
}

type NewsHandler struct {
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (h *NewsHandler) GetNews(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	bytes, err := os.ReadFile(".././example/news.json")

	if err != nil {
		return http.StatusBadRequest, err
	}

	fmt.Println(string(bytes))

	return output.SuccessResponse(w, r, &NewsResp{
		User: usr.ToClient(),
	})
}
