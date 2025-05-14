package handlers_test

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/output"
	user_repo "Arbitrax/pkg/repositories/user"
	"Arbitrax/pkg/util"

	"context"
	"net/http"
	"testing"
)

type mockUserRepo struct {
	user_repo.Repository // ✅ embed the interface (optional, for clarity/logging)
	CreateFn             func(ctx context.Context, firstName, lastName, email, password string) (*user_repo.Model, error)
	DoesEmailExistFn     func(ctx context.Context, email string) (bool, error)
}

func (m *mockUserRepo) Create(ctx context.Context, firstName, lastName, email, password string) (*user_repo.Model, error) {
	return m.CreateFn(ctx, firstName, lastName, email, password)
}

func (m *mockUserRepo) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	return m.DoesEmailExistFn(ctx, email)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := &mockUserRepo{
		DoesEmailExistFn: func(ctx context.Context, email string) (bool, error) {
			return false, nil
		},
		CreateFn: func(ctx context.Context, firstName, lastName, email, password string) (*user_repo.Model, error) {
			return &user_repo.Model{
				UUID:      "test-uuid",
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
			}, nil
		},
	}
	// inside TestRegister_Success
	handler := handlers.NewAuthHandler(mockRepo)
	wrapped := output.MakeJsonHandler(handler.Register)

	body := map[string]string{
		"first_name": "Test",
		"last_name":  "User",
		"email":      "test@example.com",
		"password":   "secure123",
	}

	res, status := util.TestJsonRequestAndDecode[handlers.ManualAuthResp](t, wrapped, http.MethodPost, "/api/auth/register", body)

	if status != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", status)
	}

	if res.User.Email != "test@example.com" {
		t.Errorf("expected email %q, got %q", "test@example.com", res.User.Email)
	}

	if res.Token == "" {
		t.Error("expected token to be set")
	}
}
