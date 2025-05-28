package handlers

import (
	"Arbitrax/pkg/constants"
	user_repo "Arbitrax/pkg/repositories/user"
	"fmt"
	"net/http"
)

func GetUserFromCtx(r *http.Request) (*user_repo.Model, error) {
	usr, ok := r.Context().Value(constants.USER_CTX).(*user_repo.Model)

	if !ok {
		return nil, fmt.Errorf("handlers.GetUserFromContext: cant find user in r.Context")
	}

	return usr, nil

}
